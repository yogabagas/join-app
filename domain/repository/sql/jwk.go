package sql

import (
	"context"
	"database/sql"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/service/jwk/repository"
	"strings"
	"time"
)

const (
	insertJWK              = "INSERT INTO jwk (id, key_text, expired_at) VALUES ($1,$2,$3)"
	updateJWK              = "UPDATE jwk SET key_text = $1, expired_at = $2 WHERE id = $3"
	selectUnexpiredKeyByID = "SELECT id, key_text, expired_at FROM jwk WHERE id = $1 AND expired_at > to_timestamp($2)"
	selectUnexpiredKey     = "SELECT id, key_text, expired_at FROM jwk WHERE expired_at > to_timestamp($1)"
)

type JWKRepositoryImpl struct {
	db DBExecutor
}

func NewJWKRepository(db DBExecutor) repository.JWKRepository {
	return &JWKRepositoryImpl{db: db}
}

func (jr *JWKRepositoryImpl) UpsertJWK(ctx context.Context, req *model.JWK) error {

	_, err := jr.db.ExecContext(ctx, insertJWK, req.ID, req.Key, req.ExpiredAt)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			_, err = jr.db.ExecContext(ctx, updateJWK, req.Key, req.ExpiredAt, req.ID)
			if err != nil {
				return err
			}
		}
		return err
	}

	return nil
}

func (jr *JWKRepositoryImpl) ReadUnexpiredKeyByID(ctx context.Context, req *model.ReadUnexpiredKeyByIDReq) (resp *model.ReadUnexpiredKeyByIDResp, err error) {

	resp = &model.ReadUnexpiredKeyByIDResp{}

	now := time.Now().Unix()

	err = jr.db.QueryRowContext(ctx, selectUnexpiredKeyByID, req.KeyID, now).Scan(&resp.ID, &resp.Key, &resp.ExpiredAt)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return resp, nil

}

func (jr *JWKRepositoryImpl) ReadUnexpiredKeys(ctx context.Context) (resp []*model.ReadUnexpiredKeyResp, err error) {

	now := time.Now().Unix()

	rows, err := jr.db.QueryContext(ctx, selectUnexpiredKey, now)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		res := &model.ReadUnexpiredKeyResp{}
		err = rows.Scan(&res.ID, &res.Key, &res.ExpiredAt)
		if err != nil {
			return nil, err
		}
		resp = append(resp, res)
	}

	return
}

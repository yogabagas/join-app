package sql

import (
	"context"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/service/jwk/repository"
)

const (
	insertJWK          = `INSERT INTO jwk (id, key, expired_at) VALUES (?,?,?)`
	selectUnexpiredKey = `SELECT id, key, expired_at FROM jwk WHERE expired_at > ?`
)

type JWKRepositoryImpl struct {
	db DBExecutor
}

func NewJWKRepository(db DBExecutor) repository.JWKRepository {
	return &JWKRepositoryImpl{db: db}
}

func (jr *JWKRepositoryImpl) CreateJWK(ctx context.Context, req *model.JWK) error {

	_, err := jr.db.ExecContext(ctx, insertJWK, req.ID, req.Key, req.ExpiredAt)
	if err != nil {
		return err
	}

	return nil
}

func (jr *JWKRepositoryImpl) ReadUnexpiredKey(ctx context.Context, req *model.ReadUnexpiredKeyReq) (resp []*model.ReadUnexpiredKeyResp, err error) {

	rows, err := jr.db.QueryContext(ctx, selectUnexpiredKey, req.Time)
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

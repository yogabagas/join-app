package sql

import (
	"context"
	"github/yogabagas/print-in/domain/model"
	"github/yogabagas/print-in/service/authz/repository"
	"time"
)

const (
	insertAuthz = `INSERT INTO authz (uid, user_uid, role_uid, last_active, created_by, updated_by)
	VALUES (?,?,?,?,?,?)`
)

type AuthzRepositoryImpl struct {
	db DBExecutor
}

func NewAuthzRepository(db DBExecutor) repository.AuthzRepository {
	return &AuthzRepositoryImpl{db: db}
}

func (ar *AuthzRepositoryImpl) CreateAuthz(ctx context.Context, req *model.Authz) error {

	now := time.Now()

	_, err := ar.db.ExecContext(ctx, insertAuthz, req.UID, req.UserUID, req.RoleUID, now, req.CreatedBy, req.UpdatedBy)
	if err != nil {
		return err
	}

	return nil

}

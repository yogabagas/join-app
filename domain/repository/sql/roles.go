package sql

import (
	"context"
	"database/sql"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/service/roles/repository"
)

const (
	insertRoles     = `INSERT INTO roles (uid, name, created_by, updated_by) VALUES (?,?,?,?)`
	selectRolesByID = `SELECT id, uid, name, is_deleted, created_by, created_at, updated_by, updated_at 
	FROM roles WHERE id = ?`
)

type RolesRepositoryImpl struct {
	db DBExecutor
}

func NewRolesRepository(db DBExecutor) repository.RolesRepository {
	return &RolesRepositoryImpl{db: db}
}

func (rr *RolesRepositoryImpl) CreateRoles(ctx context.Context, req *model.Role) error {

	_, err := rr.db.ExecContext(ctx, insertRoles, req.UID, req.Name, req.CreatedBy, req.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (rr *RolesRepositoryImpl) ReadRolesByID(ctx context.Context, req *model.ReadRolesByIDReq) (resp *model.Role, err error) {

	resp = &model.Role{}

	err = rr.db.QueryRowContext(ctx, selectRolesByID, req.ID).
		Scan(&resp.ID, &resp.UID, &resp.Name, &resp.IsDeleted, &resp.CreatedBy, &resp.CreatedAt, &resp.UpdatedBy, &resp.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return resp, nil
}

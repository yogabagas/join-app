package sql

import (
	"context"
	"database/sql"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/service/resources/repository"
	"strings"
)

const (
	insertResources = `INSERT INTO resources (uid, name, parent_uid, type, action, created_by, updated_by) VALUES (?,?,?,?,?,?,?)`
)

type ResourcesRepositoryImpl struct {
	db DBExecutor
}

func NewResourcesRepository(db DBExecutor) repository.ResourcesRepository {
	return &ResourcesRepositoryImpl{db: db}
}

func (rr *ResourcesRepositoryImpl) CreateResources(ctx context.Context, req *model.Resource) error {

	parentUID := sql.NullString{}
	if req.ParentUID != "" {
		parentUID.String = req.ParentUID
		parentUID.Valid = true
	}

	_, err := rr.db.ExecContext(ctx, insertResources, req.UID, req.Name, parentUID, req.Type, req.Action,
		req.CreatedBy, req.UpdatedBy)
	if err != nil && !strings.Contains(err.Error(), "duplicate") {
		return err
	}

	return nil
}

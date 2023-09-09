package sql

import (
	"context"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/service/modules/repository"
	"strings"
)

const (
	insertModules         = `INSERT INTO modules (uid, name, description, file, created_by ) VALUES (?,?,?,?,?)`
	insertModuleMaterials = `INSERT INTO module_materials (uid, module_uid, topic, description, created_by ) VALUES (?,?,?,?,?)`
)

type ModulesRepositoryImpl struct {
	db DBExecutor
}

func NewModulesRepository(db DBExecutor) repository.ModulesRepository {
	return &ModulesRepositoryImpl{db: db}
}

func (rr *ModulesRepositoryImpl) CreateModules(ctx context.Context, req *model.Module) error {

	_, err := rr.db.ExecContext(ctx, insertModules, req.UID, req.Name, req.Description, req.File, req.CreatedBy)
	if err != nil && !strings.Contains(err.Error(), "duplicate") {
		return err
	}

	return nil
}

func (rr *ModulesRepositoryImpl) CreateModuleMaterials(ctx context.Context, req *model.ModuleMaterial) error {

	_, err := rr.db.ExecContext(ctx, insertModuleMaterials, req.UID, req.ModuleUID, req.Topic, req.Description, req.CreatedBy)
	if err != nil && !strings.Contains(err.Error(), "duplicate") {
		return err
	}

	return nil
}

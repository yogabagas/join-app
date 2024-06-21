package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/service/modules/repository"
	"strings"
)

const (
	insertModules              = `INSERT INTO modules (uid, name, description, file, created_by ) VALUES (?,?,?,?,?)`
	updateModules              = `UPDATE modules SET name=?, description=?, file=?, updated_by=? WHERE uid=?`
	insertModuleMaterials      = `INSERT INTO module_materials (uid, module_uid, topic, description, created_by ) VALUES (?,?,?,?,?)`
	updateModuleMaterials      = `UPDATE module_materials SET module_uid= ?, topic=?, description=?, updated_by=? WHERE uid=?`
	selectModuleWithPagination = `SELECT uid, name, description, file FROM modules`
	selectCountModules         = `SELECT COUNT(*) FROM modules WHERE is_deleted = ?`
	deleteModules              = `UPDATE modules SET is_deleted = 1, updated_by=? WHERE uid=?`
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

func (rr *ModulesRepositoryImpl) UpdateModules(ctx context.Context, req *model.Module) error {

	_, err := rr.db.ExecContext(ctx, updateModules, req.Name, req.Description, req.File, req.UpdatedBy, req.UID)
	if err != nil {
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

func (rr *ModulesRepositoryImpl) UpdateModuleMaterials(ctx context.Context, req *model.ModuleMaterial) error {

	_, err := rr.db.ExecContext(ctx, updateModuleMaterials, req.ModuleUID, req.Topic, req.Description, req.UpdatedBy, req.UID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *ModulesRepositoryImpl) ReadModulesWithPagination(ctx context.Context, req *model.ReadModulesWithPaginationReq) (resp *model.ReadModulesWithPaginationResp, err error) {

	cond := fmt.Sprintf("WHERE name like %s LIMIT %d OFFSET %d", fmt.Sprint("'%"+req.Name+"%'"), req.Limit, req.Offset)

	if req.Name == "" {
		cond = fmt.Sprintf(" LIMIT %d OFFSET %d", req.Limit, req.Offset)
	}

	q := selectModuleWithPagination + cond

	rows, err := ur.db.QueryContext(ctx, q)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	resp = &model.ReadModulesWithPaginationResp{}

	for rows.Next() {
		module := model.Module{}

		err = rows.Scan(&module.UID, &module.Name, &module.Description, &module.File)
		if err != nil {
			return nil, err
		}

		resp.PerPage += 1
		resp.Modules = append(resp.Modules, module)
	}
	return resp, nil
}

func (ur *ModulesRepositoryImpl) CountModules(ctx context.Context, req *model.CountModulesReq) (resp *model.CountModulesResp, err error) {

	resp = &model.CountModulesResp{}

	err = ur.db.QueryRowContext(ctx, selectCountModules, req.IsDeleted).Scan(&resp.Total)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return
}

func (ur *ModulesRepositoryImpl) DeleteModules(ctx context.Context, uid string, userUUID string) error {
	_, err := ur.db.ExecContext(ctx, deleteModules, uid, userUUID)
	if err != nil {
		return err
	}

	return nil
}

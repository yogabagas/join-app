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
	insertModuleMaterials      = `INSERT INTO module_materials (uid, module_uid, topic, description, created_by ) VALUES (?,?,?,?,?)`
	selectModuleWithPagination = `SELECT u.uid, u.first_name, u.last_name, u.email, u.birthdate, u.username, u.created_at, 
	(SELECT COUNT(*) from users us WHERE us.id = u.id) as per_page, r.name as role_name FROM users u JOIN authz a ON u.uid = a.user_uid 
	JOIN roles r ON a.role_uid = r.uid %s`
	selectCountModules = `SELECT COUNT(*) FROM users WHERE is_deleted = ?`
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

func (ur *ModulesRepositoryImpl) ReadModulesWithPagination(ctx context.Context, req *model.ReadModulesWithPaginationReq) (resp *model.ReadModulesWithPaginationResp, err error) {

	cond := fmt.Sprintf("WHERE MATCH (u.name) AGAINST ('%s*' IN BOOLEAN MODE) LIMIT %d OFFSET %d", req.Name, req.Limit, req.Offset)

	if req.Name == "" {
		cond = fmt.Sprintf("LIMIT %d OFFSET %d", req.Limit, req.Offset)
	}

	q := fmt.Sprintf(selectModuleWithPagination, cond)

	rows, err := ur.db.QueryContext(ctx, q)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	resp = &model.ReadModulesWithPaginationResp{}

	for rows.Next() {
		module := model.Module{}
		var perPage int

		err = rows.Scan(&module.UID, &module.Name, &module.Description, &module.File)
		if err != nil {
			return nil, err
		}

		resp.PerPage += perPage
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

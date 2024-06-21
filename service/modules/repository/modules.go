package repository

import (
	"context"
	"github/yogabagas/join-app/domain/model"
)

type ModulesRepository interface {
	CreateModules(ctx context.Context, req *model.Module) error
	UpdateModules(ctx context.Context, req *model.Module) error
	CreateModuleMaterials(ctx context.Context, req *model.ModuleMaterial) error
	UpdateModuleMaterials(ctx context.Context, req *model.ModuleMaterial) error
	ReadModulesWithPagination(ctx context.Context, req *model.ReadModulesWithPaginationReq) (*model.ReadModulesWithPaginationResp, error)
	CountModules(ctx context.Context, req *model.CountModulesReq) (*model.CountModulesResp, error)
	DeleteModules(ctx context.Context, uid string, userUUID string) error
}

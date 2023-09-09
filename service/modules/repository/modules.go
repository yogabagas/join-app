package repository

import (
	"context"
	"github/yogabagas/join-app/domain/model"
)

type ModulesRepository interface {
	CreateModules(ctx context.Context, req *model.Module) error
	CreateModuleMaterials(ctx context.Context, req *model.ModuleMaterial) error
}

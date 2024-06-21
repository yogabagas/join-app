package controller

import (
	"context"
	"github/yogabagas/join-app/domain/service"
	"github/yogabagas/join-app/service/modules/usecase"
	"github/yogabagas/join-app/shared/util"
)

type ModulesControllerImpl struct {
	modulesSvc usecase.ModulesService
}

type ModulesController interface {
	CreateModules(ctx context.Context, req service.CreateModulesReq, userData *util.UserData) error
	UpdateModules(ctx context.Context, req service.UpdateModulesReq, userData *util.UserData) error
	GetModulesWithPagination(ctx context.Context, req service.GetModulesWithPaginationReq) (service.GetModulesWithPaginationResp, error)
	DeleteModules(ctx context.Context, uid string, userData util.UserData) error
}

func NewModulesController(modulesService usecase.ModulesService) ModulesController {
	return &ModulesControllerImpl{modulesSvc: modulesService}
}

func (cs *ModulesControllerImpl) CreateModules(ctx context.Context, req service.CreateModulesReq, userData *util.UserData) error {
	return cs.modulesSvc.CreateModules(ctx, req, userData)
}

func (cs *ModulesControllerImpl) UpdateModules(ctx context.Context, req service.UpdateModulesReq, userData *util.UserData) error {
	return cs.modulesSvc.UpdateModules(ctx, req, userData)
}

func (cs *ModulesControllerImpl) GetModulesWithPagination(ctx context.Context, req service.GetModulesWithPaginationReq) (service.GetModulesWithPaginationResp, error) {
	return cs.modulesSvc.GetModulesWithPagination(ctx, req)
}

func (rc *ModulesControllerImpl) DeleteModules(ctx context.Context, uid string, userData util.UserData) error {
	return rc.modulesSvc.DeleteModules(ctx, uid, userData)
}

package controller

import (
	"context"
	"github/yogabagas/join-app/domain/service"
	"github/yogabagas/join-app/service/roles/usecase"
	"strings"
)

type RolesControllerImpl struct {
	rolesSvc usecase.RolesService
}

type RolesController interface {
	CreateRoles(ctx context.Context, req service.CreateRolesReq) error
}

func NewRolesController(rolesSvc usecase.RolesService) RolesController {
	return &RolesControllerImpl{rolesSvc: rolesSvc}
}

func (rc *RolesControllerImpl) CreateRoles(ctx context.Context, req service.CreateRolesReq) error {

	req.Name = strings.ToLower(req.Name)

	return rc.rolesSvc.CreateRoles(ctx, req)
}

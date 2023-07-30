package usecase

import (
	"context"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/domain/service"
	"github/yogabagas/join-app/service/roles/repository"
	"github/yogabagas/join-app/shared/util"
)

type RolesServiceImpl struct {
	rolesRepo repository.RolesRepository
}

type RolesService interface {
	CreateRoles(ctx context.Context, req service.CreateRolesReq) error
}

func NewRolesService(rolesRepo repository.RolesRepository) RolesService {
	return &RolesServiceImpl{rolesRepo: rolesRepo}
}

func (rs *RolesServiceImpl) CreateRoles(ctx context.Context, req service.CreateRolesReq) error {

	uID := util.NewULIDGenerate()

	return rs.rolesRepo.CreateRoles(ctx, &model.Role{
		UID:       uID,
		Name:      req.Name,
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.CreatedBy,
	})
}

package usecase

import (
	"context"
	"github/yogabagas/print-in/domain/model"
	"github/yogabagas/print-in/domain/service"
	"github/yogabagas/print-in/service/roles/repository"
	"github/yogabagas/print-in/shared/util"
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

	err := rs.rolesRepo.CreateRoles(ctx, &model.Role{
		UID:       uID,
		Name:      req.Name,
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.CreatedBy,
	})
	if err != nil {
		return err
	}

	return nil
}

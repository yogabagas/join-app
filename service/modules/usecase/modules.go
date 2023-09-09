package usecase

import (
	"context"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/domain/service"
	coursesRepo "github/yogabagas/join-app/service/modules/repository"
	"github/yogabagas/join-app/shared/util"
)

type ModulesServiceImpl struct {
	modulesRepo coursesRepo.ModulesRepository
}

type ModulesService interface {
	CreateModules(ctx context.Context, req service.CreateModulesReq, userData *util.UserData) error
}

func NewModulesService(modulesRepo coursesRepo.ModulesRepository) ModulesService {
	return &ModulesServiceImpl{
		modulesRepo: modulesRepo}
}

func (cs *ModulesServiceImpl) CreateModules(ctx context.Context, req service.CreateModulesReq, userData *util.UserData) error {

	uID := util.NewULIDGenerate()
	err := cs.modulesRepo.CreateModules(ctx, &model.Module{
		UID:         uID,
		Name:        req.Name,
		Description: req.Description,
		File:        req.File,
		CreatedBy:   userData.UserUUID,
	})

	for _, moduleMaterial := range req.ModuleMaterial {

		uIDModuleMaterial := util.NewULIDGenerate()
		_ = cs.modulesRepo.CreateModuleMaterials(ctx, &model.ModuleMaterial{
			UID:         uIDModuleMaterial,
			ModuleUID:   uID,
			Topic:       moduleMaterial.Topic,
			Description: moduleMaterial.Description,
			CreatedBy:   userData.UserUUID,
		})

	}

	return err
}

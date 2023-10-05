package usecase

import (
	"context"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/domain/repository/sql"
	"github/yogabagas/join-app/domain/service"
	"github/yogabagas/join-app/service/modules/presenter"
	coursesRepo "github/yogabagas/join-app/service/modules/repository"
	"github/yogabagas/join-app/shared/constant"
	"github/yogabagas/join-app/shared/util"
)

type ModulesServiceImpl struct {
	modulesRepo coursesRepo.ModulesRepository
	presenter   presenter.ModulesPresenter
}

type ModulesService interface {
	CreateModules(ctx context.Context, req service.CreateModulesReq, userData *util.UserData) error
	UpdateModules(ctx context.Context, req service.UpdateModulesReq, userData *util.UserData) error
	GetModulesWithPagination(ctx context.Context, req service.GetModulesWithPaginationReq) (service.GetModulesWithPaginationResp, error)
	DeleteModules(ctx context.Context, uid string, userData util.UserData) error
}

func NewModulesService(repository sql.RepositoryRegistry, presenter presenter.ModulesPresenter) ModulesService {
	return &ModulesServiceImpl{
		modulesRepo: repository.ModulesRepository(),
		presenter:   presenter}
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

func (cs *ModulesServiceImpl) UpdateModules(ctx context.Context, req service.UpdateModulesReq, userData *util.UserData) error {

	err := cs.modulesRepo.UpdateModules(ctx, &model.Module{
		UID:         req.UID,
		Name:        req.Name,
		Description: req.Description,
		File:        req.File,
		UpdatedBy:   userData.UserUUID,
	})

	for _, moduleMaterial := range req.ModuleMaterial {

		_ = cs.modulesRepo.UpdateModuleMaterials(ctx, &model.ModuleMaterial{
			UID:         moduleMaterial.UID,
			ModuleUID:   req.UID,
			Topic:       moduleMaterial.Topic,
			Description: moduleMaterial.Description,
			UpdatedBy:   userData.UserUUID,
		})

	}

	return err
}

func (cs *ModulesServiceImpl) GetModulesWithPagination(ctx context.Context, req service.GetModulesWithPaginationReq) (resp service.GetModulesWithPaginationResp, err error) {

	modules, err := cs.modulesRepo.ReadModulesWithPagination(ctx, &model.ReadModulesWithPaginationReq{
		UID:    req.UID,
		Name:   req.Name,
		Limit:  req.Limit,
		Offset: util.PageToOffset(req.Limit, req.Page),
	})
	if err != nil {
		return
	}

	count, err := cs.modulesRepo.CountModules(ctx, &model.CountModulesReq{
		IsDeleted: constant.False.Int(),
	})
	if err != nil {
		return
	}

	return cs.presenter.GetModulesWithPagination(ctx, req, modules, count)
}

func (rs *ModulesServiceImpl) DeleteModules(ctx context.Context, uid string, userData util.UserData) error {
	return rs.modulesRepo.DeleteModules(ctx, uid, userData.UserUUID)
}

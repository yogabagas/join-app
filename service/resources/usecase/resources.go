package usecase

import (
	"context"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/domain/repository/sql"
	"github/yogabagas/join-app/domain/service"
	resourcesRepo "github/yogabagas/join-app/service/resources/repository"
	"github/yogabagas/join-app/shared/util"
)

type ResourcesServiceImpl struct {
	resourcesRepo resourcesRepo.ResourcesRepository
}

type ResourcesService interface {
	CreateResources(ctx context.Context, req service.CreateResourcesReq) error
}

func NewResourcesService(repository sql.RepositoryRegistry) ResourcesService {
	return &ResourcesServiceImpl{
		resourcesRepo: repository.ResourcesRepository()}
}

func (rs *ResourcesServiceImpl) CreateResources(ctx context.Context, req service.CreateResourcesReq) error {

	uID := util.NewULIDGenerate()

	return rs.resourcesRepo.CreateResources(ctx, &model.Resource{
		UID:       uID,
		Name:      req.Name,
		Type:      req.Type,
		Action:    req.Action,
		ParentUID: req.ParentUID,
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.CreatedBy,
	})
}

package registry

import (
	"github/yogabagas/join-app/adapter/controller"
	"github/yogabagas/join-app/domain/repository/sql"
	"github/yogabagas/join-app/service/resources/repository"
	"github/yogabagas/join-app/service/resources/usecase"
)

func (m *module) NewResourcesRepository() repository.ResourcesRepository {
	return sql.NewResourcesRepository(m.sqlDB)
}

func (m *module) NewResourcesRegistry() usecase.ResourcesService {
	return usecase.NewResourcesService(m.NewRepositoryRegistry())
}

func (m *module) NewResourcesController() controller.ResourcesController {
	return controller.NewResourcesController(m.NewResourcesRegistry())
}

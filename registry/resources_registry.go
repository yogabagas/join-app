package registry

import (
	"github/yogabagas/join-app/service/resources/presenter"
	"github/yogabagas/join-app/service/resources/usecase"
)

func (m *module) NewResourcesRepository() repository.ResourcesRepository {
	return sql.NewResourcesRepository(m.sqlDB)
}

func (m *module) NewResourcesService() usecase.ResourcesService {
	return usecase.NewResourcesService(
		m.NewCacheRegistry(),
		m.NewRepositoryRegistry(),
		m.NewResourcesPresenter())
}

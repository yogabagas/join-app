package registry

import (
	"github/yogabagas/join-app/adapter/controller"
	"github/yogabagas/join-app/service/authz/presenter"
	"github/yogabagas/join-app/service/authz/usecase"
)

func (m *module) NewAuthzPresenter() presenter.AuthzPresenter {
	return presenter.NewAuthzPresenter()
}

func (m *module) NewAuthzRegistry() usecase.AuthzService {
	return usecase.NewAuthzService(
		m.NewRepositoryRegistry(),
		m.NewCacheRegistry(),
		m.NewAuthzPresenter(),
	)
}

func (m *module) NewAuthzController() controller.AuthzController {
	return controller.NewAuthzController(m.NewAuthzRegistry())
}

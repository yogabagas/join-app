package registry

import (
	"github/yogabagas/join-app/adapter/controller"
	"github/yogabagas/join-app/service/modules/presenter"
	"github/yogabagas/join-app/service/modules/usecase"
)

func (m *module) NewModulesPresenter() presenter.ModulesPresenter {
	return presenter.NewModulesPresenter()
}

func (m *module) NewModulesRegistry() usecase.ModulesService {
	return usecase.NewModulesService(m.NewRepositoryRegistry(), m.NewModulesPresenter())
}

func (m *module) NewCoursesController() controller.ModulesController {
	return controller.NewModulesController(m.NewModulesRegistry())
}

package registry

import (
	"github/yogabagas/join-app/adapter/controller"
	"github/yogabagas/join-app/domain/repository/sql"
	"github/yogabagas/join-app/service/modules/repository"
	"github/yogabagas/join-app/service/modules/usecase"
)

func (m *module) NewCoursesRepository() repository.ModulesRepository {
	return sql.NewModulesRepository(m.sqlDB)
}

func (m *module) NewCoursesRegistry() usecase.ModulesService {
	return usecase.NewModulesService(m.NewCoursesRepository())
}

func (m *module) NewCoursesController() controller.ModulesController {
	return controller.NewModulesController(m.NewCoursesRegistry())
}

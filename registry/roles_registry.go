package registry

import (
	"github/yogabagas/print-in/adapter/controller"
	"github/yogabagas/print-in/domain/repository/sql"
	"github/yogabagas/print-in/service/roles/repository"
	"github/yogabagas/print-in/service/roles/usecase"
)

func (m *module) NewRolesRepository() repository.RolesRepository {
	return sql.NewRolesRepository(m.sqlDB)
}

func (m *module) NewRolesRegistry() usecase.RolesService {
	return usecase.NewRolesService(m.NewRolesRepository())
}

func (m *module) NewRolesController() controller.RolesController {
	return controller.NewRolesController(m.NewRolesRegistry())
}

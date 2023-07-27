package registry

import (
	"github/yogabagas/join-app/adapter/controller"
	"github/yogabagas/join-app/domain/repository/sql"
	"github/yogabagas/join-app/service/roles/repository"
	"github/yogabagas/join-app/service/roles/usecase"
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

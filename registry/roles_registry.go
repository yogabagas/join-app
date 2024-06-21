package registry

import (
	"github/yogabagas/join-app/service/roles/usecase"
)

func (m *module) NewRolesController() controller.RolesController {
	return controller.NewRolesController(m.NewRolesRegistry())
func (m *module) NewRolesService() usecase.RolesService {
	return usecase.NewRolesService(m.NewRepositoryRegistry())
}

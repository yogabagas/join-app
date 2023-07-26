package registry

import (
	"github/yogabagas/print-in/adapter/controller"
	"github/yogabagas/print-in/service/users/usecase"
)

func (m *module) NewUsersRegistry() usecase.UsersService {
	return usecase.NewUsersService(m.NewRepositoryRegistry(), m.NewSessionRepositoryRegistry())
}

func (m *module) NewUsersController() controller.UsersController {
	return controller.NewUsersController(m.NewUsersRegistry())
}

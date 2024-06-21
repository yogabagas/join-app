package registry

import (
	"github/yogabagas/join-app/service/authz/usecase"
)

func (m *module) NewAuthzService() usecase.AuthzService {
	return usecase.NewAuthzService(
		m.NewRepositoryRegistry(),
		m.NewCacheRegistry(),
	)
}

// func (m *module) NewAuthzController() controller.AuthzController {
// 	return controller.NewAuthzController(m.NewAuthzRegistry())
// }

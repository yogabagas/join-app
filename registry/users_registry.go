package registry

import (
	"github/yogabagas/join-app/service/users/presenter"
	"github/yogabagas/join-app/service/users/usecase"
)

func (m *module) NewUsersPresenter() presenter.UsersPresenter {
	return presenter.NewUsersPresenter()
}

func (m *module) NewUsersService() usecase.UsersService {
	return usecase.NewUsersService(
		m.NewRepositoryRegistry(),
		m.NewCacheRegistry(),
		m.NewUsersPresenter())
}

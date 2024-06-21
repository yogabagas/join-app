package registry

import (
	"github/yogabagas/join-app/service/access/presenter"
	"github/yogabagas/join-app/service/access/usecase"
)

func (m *module) NewAccessPresenter() presenter.AccessPresenter {
	return presenter.NewAccessPresenter()
}
func (m *module) NewAccessService() usecase.AccessService {
	return usecase.NewAccessService(
		m.NewRepositoryRegistry(),
		m.NewCacheRegistry(),
		m.NewAccessPresenter(),
	)
}

package registry

import (
	"github/yogabagas/join-app/service/jwk/presenter"
	"github/yogabagas/join-app/service/jwk/usecase"
)

func (m *module) NewJWKPresenter() presenter.JWKPresenter {
	return presenter.NewJWKPresenter()
}

func (m *module) NewJWKService() usecase.JWKService {
	return usecase.NewJWKService(
		m.NewRepositoryRegistry(),
		m.NewCacheRegistry(),
		m.NewJWKPresenter(),
	)
}

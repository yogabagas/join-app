package registry

import (
	"database/sql"
	"github/yogabagas/join-app/domain/repository/cache"
	repo "github/yogabagas/join-app/domain/repository/sql"
	"github/yogabagas/join-app/service"

	"github.com/go-redis/redis/v8"
)

type module struct {
	sqlDB *sql.DB
	cache cache.Cache
}

type Registry interface {
	NewAppService() service.ServiceRegistry
}

type Option func(*module)

func NewRegistry(opts ...Option) Registry {
	m := &module{}

	for _, o := range opts {
		o(m)
	}

	return m
}

func NewSQLConn(db *sql.DB) Option {
	return func(m *module) {
		m.sqlDB = db
	}
}

func NewCache(cache cache.Cache) Option {
	return func(m *module) {
		m.cache = cache
	}
}

func (m *module) NewRepositoryRegistry() repo.RepositoryRegistry {
	return repo.NewRepositoryRegistry(m.sqlDB, m.cache)
}

func (m *module) NewSessionRepositoryRegistry() repoCache.RepositoryRegistry {
	return repoCache.NewRepositoryRegistry(m.sqlDB, m.cache)
}

func (m *module) NewAppService() service.ServiceRegistry {
	return service.ServiceRegistry{
		AccessService:    m.NewAccessService(),
		AuthzService:     m.NewAuthzService(),
		JwkService:       m.NewJWKService(),
		ResourcesService: m.NewResourcesService(),
		RolesService:     m.NewRolesService(),
		UsersService:     m.NewUsersService(),
	}
}

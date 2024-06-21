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
	cache *redis.Client
	ns    string
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

func NewCache(redisClient *redis.Client) Option {
	return func(m *module) {
		m.cache = redisClient
	}
}

func (m *module) NewRepositoryRegistry() repo.RepositoryRegistry {
	return repo.NewRepositoryRegistry(m.sqlDB)
}

func (m *module) NewCacheRegistry() cache.Cache {
	return cache.NewCacheRepository(m.cache, m.ns)
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

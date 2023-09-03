package registry

import (
	"database/sql"
	"github/yogabagas/join-app/adapter/controller"
	repoCache "github/yogabagas/join-app/domain/repository/cache"
	repo "github/yogabagas/join-app/domain/repository/sql"
	"github/yogabagas/join-app/pkg/cache"
)

type module struct {
	sqlDB *sql.DB
	cache cache.Cache
}

type Registry interface {
	NewAppController() controller.AppController
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

func (m *module) NewAppController() controller.AppController {
	return controller.AppController{
		UsersController:     m.NewUsersController(),
		RolesController:     m.NewRolesController(),
		ResourcesController: m.NewResourcesController(),
		CoursesController:   m.NewCoursesController(),
	}
}

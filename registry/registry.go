package registry

import (
	"database/sql"
	"github/yogabagas/join-app/adapter/controller"
	repo "github/yogabagas/join-app/domain/repository/sql"
)

type module struct {
	sqlDB *sql.DB
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

func (m *module) NewRepositoryRegistry() repo.RepositoryRegistry {
	return repo.NewRepositoryRegistry(m.sqlDB)
}

func (m *module) NewAppController() controller.AppController {
	return controller.AppController{
		UsersController: m.NewUsersController(),
		RolesController: m.NewRolesController(),
	}
}

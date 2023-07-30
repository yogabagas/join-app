package cache

import (
	"database/sql"
	"github/yogabagas/join-app/pkg/cache"
	usersRepo "github/yogabagas/join-app/service/users/repository"
)

type RepositoryRegistryImpl struct {
	db    *sql.DB
	cache cache.Cache
}

type RepositoryRegistry interface {
	SessionRepository() usersRepo.SessionRepository
}

func NewRepositoryRegistry(db *sql.DB, cache cache.Cache) RepositoryRegistry {
	return &RepositoryRegistryImpl{db: db, cache: cache}
}

func (r RepositoryRegistryImpl) SessionRepository() usersRepo.SessionRepository {
	return NewSessionRepository(r.cache)
}

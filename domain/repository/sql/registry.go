package sql

import (
	"context"
	"database/sql"
	"github/yogabagas/join-app/pkg/cache"
	authzRepo "github/yogabagas/join-app/service/authz/repository"
	rolesRepo "github/yogabagas/join-app/service/roles/repository"
	usersRepo "github/yogabagas/join-app/service/users/repository"
)

type InTransaction func(RepositoryRegistry) (interface{}, error)

type RepositoryRegistryImpl struct {
	db         *sql.DB
	cache      cache.Cache
	dbExecutor DBExecutor
}

type RepositoryRegistry interface {
	AuthzRepository() authzRepo.AuthzRepository
	RolesRepository() rolesRepo.RolesRepository
	UserRepository() usersRepo.UsersRepository

	DoInTransaction(ctx context.Context, txFunc InTransaction) (out interface{}, err error)
}

func NewRepositoryRegistry(db *sql.DB, cache cache.Cache) RepositoryRegistry {
	return &RepositoryRegistryImpl{db: db, cache: cache}
}

func (r RepositoryRegistryImpl) AuthzRepository() authzRepo.AuthzRepository {
	if r.dbExecutor != nil {
		return NewAuthzRepository(r.dbExecutor)
	}
	return NewAuthzRepository(r.db)
}

func (r RepositoryRegistryImpl) RolesRepository() rolesRepo.RolesRepository {
	if r.dbExecutor != nil {
		return NewRolesRepository(r.dbExecutor)
	}
	return NewRolesRepository(r.db)
}

func (r RepositoryRegistryImpl) UserRepository() usersRepo.UsersRepository {
	if r.dbExecutor != nil {
		return NewUsersRepository(r.dbExecutor, r.cache)
	}
	return NewUsersRepository(r.db, r.cache)
}

func (r RepositoryRegistryImpl) DoInTransaction(ctx context.Context, txFunc InTransaction) (out interface{}, err error) {
	var tx *sql.Tx

	registry := r

	if r.dbExecutor == nil {
		tx, err = r.db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}

		defer func() {
			if p := recover(); p != nil {
				_ = tx.Rollback()
				panic(p)
			} else if err != nil {
				rErr := tx.Rollback()
				if rErr != nil {
					err = rErr
				}
			} else {
				err = tx.Commit()
			}
		}()
		registry = RepositoryRegistryImpl{
			db:         r.db,
			dbExecutor: tx,
		}
	}
	out, err = txFunc(registry)
	return
}

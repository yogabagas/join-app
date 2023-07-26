package cache

import (
	"context"
	"github/yogabagas/print-in/pkg/cache"
	"github/yogabagas/print-in/service/users/repository"
)

type UsersRepositoryImpl struct {
	cache cache.Cache
}

func NewSessionRepository(cache cache.Cache) repository.SessionRepository {
	return &UsersRepositoryImpl{cache: cache}
}

func (ur *UsersRepositoryImpl) CreateSession(ctx context.Context, userUUID string) error {
	return ur.cache.Set(ctx, "user_uuid:"+userUUID, true, 86400)
}

func (ur *UsersRepositoryImpl) DeleteSession(ctx context.Context, userUUID string) error {
	return ur.cache.Delete(ctx, "user_uuid:"+userUUID)
}

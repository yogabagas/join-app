package repository

import (
	"context"
	"github/yogabagas/print-in/domain/model"
)

type UsersRepository interface {
	CreateUsers(ctx context.Context, req *model.User) error
	ReadUserByEmail(ctx context.Context, email string) (*model.Session, error)
	CreateSession(ctx context.Context, userUUID string) error
	DeleteSession(ctx context.Context, userUUID string) error
}

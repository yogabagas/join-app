package repository

import (
	"context"
	"github/yogabagas/print-in/domain/model"
)

type UsersRepository interface {
	CreateUsers(ctx context.Context, req *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

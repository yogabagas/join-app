package repository

import (
	"context"
	"github/yogabagas/join-app/domain/model"
)

type AuthzRepository interface {
	CreateAuthz(ctx context.Context, req *model.Authz) error
}

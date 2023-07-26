package repository

import (
	"context"
	"github/yogabagas/print-in/domain/model"
)

type AuthzRepository interface {
	CreateAuthz(ctx context.Context, req *model.Authz) error
}

package repository

import (
	"context"
	"github/yogabagas/join-app/domain/model"
)

type RolesRepository interface {
	CreateRoles(ctx context.Context, req *model.Role) error
	ReadRolesByID(ctx context.Context, req *model.ReadRolesByIDReq) (*model.Role, error)
}

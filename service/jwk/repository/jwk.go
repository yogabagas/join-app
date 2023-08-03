package repository

import (
	"context"
	"github/yogabagas/join-app/domain/model"
)

type JWKRepository interface {
	CreateJWK(context.Context, *model.JWK) error
	ReadUnexpiredKey(context.Context, *model.ReadUnexpiredKeyReq) ([]*model.ReadUnexpiredKeyResp, error)
}

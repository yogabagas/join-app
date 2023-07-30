package repository

import (
	"context"
	"github/yogabagas/join-app/domain/model"
)

type ResourcesRepository interface {
	CreateResources(ctx context.Context, req *model.Resource) error
}

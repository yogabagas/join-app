package controller

import (
	"context"
	"github/yogabagas/print-in/domain/service"
	"github/yogabagas/print-in/service/users/usecase"
)

type UsersControllerImpl struct {
	usersSvc usecase.UsersService
}

type UsersController interface {
	CreateUsers(ctx context.Context, req service.CreateUsersReq) error
}

func NewUsersController(userSvc usecase.UsersService) UsersController {
	return &UsersControllerImpl{usersSvc: userSvc}
}

func (uc *UsersControllerImpl) CreateUsers(ctx context.Context, req service.CreateUsersReq) error {
	return uc.usersSvc.CreateUsers(ctx, req)
}

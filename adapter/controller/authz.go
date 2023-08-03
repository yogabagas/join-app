package controller

import (
	"context"
	"github/yogabagas/join-app/domain/service"
	"github/yogabagas/join-app/service/authz/usecase"
)

type AuthzControllerImpl struct {
	authzSvc usecase.AuthzService
}

type AuthzController interface {
	Login(ctx context.Context, req service.LoginReq) (resp service.LoginResp, err error)
}

func NewAuthzController(authzSvc usecase.AuthzService) AuthzController {
	return &AuthzControllerImpl{
		authzSvc: authzSvc,
	}
}

func (ac *AuthzControllerImpl) Login(ctx context.Context, req service.LoginReq) (resp service.LoginResp, err error) {
	return ac.authzSvc.Login(ctx, req)
}

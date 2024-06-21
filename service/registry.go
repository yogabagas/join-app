package service

import (
	accessSvc "github/yogabagas/join-app/service/access/usecase"
	authzSvc "github/yogabagas/join-app/service/authz/usecase"
	jwkSvc "github/yogabagas/join-app/service/jwk/usecase"
	resourcesSvc "github/yogabagas/join-app/service/resources/usecase"
	rolesSvc "github/yogabagas/join-app/service/roles/usecase"
	usersSvc "github/yogabagas/join-app/service/users/usecase"
)

type ServiceRegistry struct {
	AccessService    interface{ accessSvc.AccessService }
	AuthzService     interface{ authzSvc.AuthzService }
	JwkService       interface{ jwkSvc.JWKService }
	ResourcesService interface{ resourcesSvc.ResourcesService }
	RolesService     interface{ rolesSvc.RolesService }
	UsersService     interface{ usersSvc.UsersService }
}

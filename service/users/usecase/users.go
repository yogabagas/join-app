package usecase

import (
	"context"
	"errors"
	"github/yogabagas/join-app/config"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/domain/repository/sql"
	"github/yogabagas/join-app/domain/service"
	authzRepo "github/yogabagas/join-app/service/authz/repository"
	rolesRepo "github/yogabagas/join-app/service/roles/repository"
	"github/yogabagas/join-app/service/users/presenter"
	usersRepo "github/yogabagas/join-app/service/users/repository"

	"github/yogabagas/join-app/shared/constant"
	"github/yogabagas/join-app/shared/util"
	"log"
	"time"
)

type UsersServiceImpl struct {
	authzRepo authzRepo.AuthzRepository
	rolesRepo rolesRepo.RolesRepository
	usersRepo usersRepo.UsersRepository
	presenter presenter.UsersPresenter
}

type UsersService interface {
	CreateUsers(ctx context.Context, req service.CreateUsersReq) error
	GetUsersWithPagination(ctx context.Context, req service.GetUsersWithPaginationReq) (service.GetUsersWithPaginationResp, error)
}

func NewUsersService(repository sql.RepositoryRegistry, presenter presenter.UsersPresenter) UsersService {
	return &UsersServiceImpl{
		authzRepo: repository.AuthzRepository(),
		rolesRepo: repository.RolesRepository(),
		usersRepo: repository.UserRepository(),
		presenter: presenter,
	}
}

func (us *UsersServiceImpl) CreateUsers(ctx context.Context, req service.CreateUsersReq) error {

	role, err := us.rolesRepo.ReadRolesByID(ctx, &model.ReadRolesByIDReq{
		ID: req.RoleID,
	})
	if err != nil {
		return err
	} else if role == nil {
		return errors.New("role is not found")
	}

	pwd, err := util.Hash(config.GlobalCfg.PasswordAlg, req.Password)
	if err != nil {
		log.Println("err hashing", err)
		return err
	}

	hbd, err := time.Parse(time.DateOnly, req.Birthdate)
	if err != nil {
		log.Println("err time parse", err)
		return err
	}

	userUID := util.NewULIDGenerate()

	err = us.usersRepo.CreateUsers(ctx, &model.User{
		UID:       userUID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Birthdate: hbd,
		Email:     req.Email,
		Username:  req.Username,
		Password:  util.Base64(pwd),
		CreatedBy: userUID,
		UpdatedBy: userUID,
	})
	if err != nil {
		log.Println("err insert data", err)
		return err
	}

	authzUID := util.NewULIDGenerate()

	err = us.authzRepo.CreateAuthz(ctx, &model.Authz{
		UID:       authzUID,
		UserUID:   userUID,
		RoleUID:   role.UID,
		CreatedBy: userUID,
		UpdatedBy: userUID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (us *UsersServiceImpl) GetUsersWithPagination(ctx context.Context, req service.GetUsersWithPaginationReq) (resp service.GetUsersWithPaginationResp, err error) {

	users, err := us.usersRepo.ReadUsersWithPagination(ctx, &model.ReadUsersWithPaginationReq{
		Fullname: req.Fullname,
		Limit:    req.Limit,
		Offset:   util.PageToOffset(req.Limit, req.Page),
	})
	if err != nil {
		return
	}

	count, err := us.usersRepo.CountUsers(ctx, &model.CountUsersReq{
		IsDeleted: constant.False.Int(),
	})
	if err != nil {
		return
	}

	return us.presenter.GetUsersWithPagination(ctx, req, users, count)
}

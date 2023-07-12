package usecase

import (
	"context"
	"errors"
	"github/yogabagas/print-in/config"
	"github/yogabagas/print-in/domain/model"
	"github/yogabagas/print-in/domain/repository/sql"
	"github/yogabagas/print-in/domain/service"
	authzRepo "github/yogabagas/print-in/service/authz/repository"
	rolesRepo "github/yogabagas/print-in/service/roles/repository"
	usersRepo "github/yogabagas/print-in/service/users/repository"

	"github/yogabagas/print-in/shared/util"
	"log"
	"time"
)

type UsersServiceImpl struct {
	authzRepo authzRepo.AuthzRepository
	rolesRepo rolesRepo.RolesRepository
	usersRepo usersRepo.UsersRepository
}

type UsersService interface {
	CreateUsers(ctx context.Context, req service.CreateUsersReq) error
}

func NewUsersService(repository sql.RepositoryRegistry) UsersService {
	return &UsersServiceImpl{
		authzRepo: repository.AuthzRepository(),
		rolesRepo: repository.RolesRepository(),
		usersRepo: repository.UserRepository(),
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

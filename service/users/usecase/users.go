package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github/yogabagas/join-app/config"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/domain/repository/cache"
	"github/yogabagas/join-app/domain/repository/sql"
	"github/yogabagas/join-app/domain/service"
	authzRepo "github/yogabagas/join-app/service/authz/repository"
	rolesRepo "github/yogabagas/join-app/service/roles/repository"
	"github/yogabagas/join-app/service/users/presenter"
	usersRepo "github/yogabagas/join-app/service/users/repository"
	"github/yogabagas/join-app/shared/util"

	"github/yogabagas/join-app/shared/constant"
	"log"
	"time"
)

type UsersServiceImpl struct {
	authzRepo   authzRepo.AuthzRepository
	rolesRepo   rolesRepo.RolesRepository
	usersRepo   usersRepo.UsersRepository
	sessionRepo usersRepo.SessionRepository
	presenter   presenter.UsersPresenter
}

type UsersService interface {
	CreateUsers(ctx context.Context, req service.CreateUsersReq) error
	Login(ctx context.Context, req service.LoginReq) (*service.LoginRes, error)
	Logout(ctx context.Context, userUUID string) (bool, error)
	GetUsersWithPagination(ctx context.Context, req service.GetUsersWithPaginationReq) (service.GetUsersWithPaginationResp, error)
}

func NewUsersService(repository sql.RepositoryRegistry, sessionRepository cache.RepositoryRegistry, presenter presenter.UsersPresenter) UsersService {
	return &UsersServiceImpl{
		authzRepo:   repository.AuthzRepository(),
		rolesRepo:   repository.RolesRepository(),
		usersRepo:   repository.UserRepository(),
		sessionRepo: sessionRepository.SessionRepository(),
		presenter:   presenter,
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

func (us *UsersServiceImpl) Login(ctx context.Context, req service.LoginReq) (*service.LoginRes, error) {
	user, err := us.usersRepo.ReadUserByEmail(ctx, req.Email)

	pwd, err := util.Hash(config.GlobalCfg.PasswordAlg, req.Password)

	if util.Base64(pwd) != user.Password {
		return nil, errors.New("Invalid Password")
	}

	accessToken, err := us.CreateToken(24, user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := us.CreateToken(24, user)
	if err != nil {
		return nil, err
	}

	err = us.sessionRepo.CreateSession(ctx, user.UserUID)
	if err != nil {
		return nil, err
	}

	data := service.LoginRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &data, err
}

func (us *UsersServiceImpl) Logout(ctx context.Context, userUUID string) (bool, error) {
	err := us.sessionRepo.DeleteSession(ctx, userUUID)
	if err != nil {
		return false, err
	}

	return true, err
}

func (us *UsersServiceImpl) CreateToken(ttl time.Duration, user *model.Session) (string, error) {
	now := time.Now().UTC()
	claims := make(jwt.MapClaims)
	claims["user_uuid"] = user.UserUID
	claims["role_uuid"] = user.RoleUID
	claims["exp"] = now.Add(ttl * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.GlobalCfg.App.JwtSecret))

	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return tokenString, nil
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

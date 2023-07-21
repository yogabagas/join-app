package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
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
	Login(ctx context.Context, req service.LoginReq) (*service.LoginRes, error)
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

func (us *UsersServiceImpl) Login(ctx context.Context, req service.LoginReq) (*service.LoginRes, error) {
	user, err := us.usersRepo.FindByEmail(ctx, req.Email)

	pwd, err := util.Hash(config.GlobalCfg.PasswordAlg, req.Password)

	if util.Base64(pwd) != user.Password {
		return nil, errors.New("Invalid Password")
	}

	tokenClaim := model.TokenClaim{}
	tokenClaim.ID = user.ID
	tokenClaim.StandardClaims.IssuedAt = time.Now().Unix()
	tokenClaim.StandardClaims.Id = util.NewULIDGenerate()

	token := createToken(user)
	birthdate := user.Birthdate.Format("2006-01-02")
	data := service.LoginRes{
		Account: service.LoginUsersRes{
			FirstName: &user.FirstName,
			LastName:  &user.LastName,
			Birthdate: &birthdate,
			Email:     &user.Email,
			Username:  &user.Username,
		},
		AccessToken: token,
	}

	return &data, err
}

func createToken(user *model.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

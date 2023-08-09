package usecase

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github/yogabagas/join-app/config"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/domain/repository/cache"
	"github/yogabagas/join-app/domain/repository/sql"
	"github/yogabagas/join-app/domain/service"
	authzRepo "github/yogabagas/join-app/service/authz/repository"
	jwkRepo "github/yogabagas/join-app/service/jwk/repository"
	rolesRepo "github/yogabagas/join-app/service/roles/repository"
	"github/yogabagas/join-app/service/users/presenter"
	usersRepo "github/yogabagas/join-app/service/users/repository"
	"github/yogabagas/join-app/shared/util"

	"github.com/go-jose/go-jose/v3"
	"github.com/golang-jwt/jwt"

	"github/yogabagas/join-app/shared/constant"
	"log"
	"time"
)

var (
	keyMap map[string]jose.Signer
)

type UsersServiceImpl struct {
	authzRepo authzRepo.AuthzRepository
	jwkRepo   jwkRepo.JWKRepository
	rolesRepo rolesRepo.RolesRepository
	usersRepo usersRepo.UsersRepository
	cache     cache.Cache
	presenter presenter.UsersPresenter
}

type UsersService interface {
	CreateUsers(ctx context.Context, req service.CreateUsersReq) error
	Logout(ctx context.Context, req service.LogoutReq) error
	GetUsersWithPagination(ctx context.Context, req service.GetUsersWithPaginationReq) (service.GetUsersWithPaginationResp, error)
}

func NewUsersService(repository sql.RepositoryRegistry, cache cache.Cache, presenter presenter.UsersPresenter) UsersService {
	return &UsersServiceImpl{
		authzRepo: repository.AuthzRepository(),
		jwkRepo:   repository.JWKRepository(),
		rolesRepo: repository.RolesRepository(),
		usersRepo: repository.UserRepository(),
		cache:     cache,
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

func (us *UsersServiceImpl) Logout(ctx context.Context, req service.LogoutReq) error {

	key := fmt.Sprintf("auth::user-uid:%s", req.UserUID)

	return us.cache.Delete(ctx, key)
}

func (us *UsersServiceImpl) generateAndSignAccessToken(ctx context.Context, req *model.GenerateAccessTokenReq) (resp *model.GenerateAccessTokenResp, err error) {

	publicKeyCache := "jwk::public-key:*"

	remaining := us.cache.RemainingTime(ctx, publicKeyCache)

	log.Printf("remaining key access token %s is %d", publicKeyCache, remaining)

	if remaining <= 0 {
		err := us.generateJWKKey(ctx)
		if err != nil {
			return nil, err
		}
	}

	signer, ok := keyMap[config.GlobalCfg.JWK.KeyID]
	if !ok {
		err := us.generateJWKKey(ctx)
		if err != nil {
			return nil, err
		}
		signer = keyMap[config.GlobalCfg.JWK.KeyID]
	}

	claims := make(jwt.MapClaims)
	claims["sub"] = req.UserUID
	claims["role_uid"] = req.RoleUID
	claims["iat"] = time.Now().UTC().Unix()
	claims["exp"] = time.Now().UTC().Add(time.Duration(req.ExpiredAt) * time.Second).Unix()
	claims["last_active"] = req.LastActive

	data, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	obj, err := signer.Sign(data)
	if err != nil {
		return nil, err
	}

	tokenString, err := obj.CompactSerialize()
	if err != nil {
		return nil, err
	}

	return &model.GenerateAccessTokenResp{
		Token: tokenString,
	}, nil
}

func (us *UsersServiceImpl) generateAndSignRefreshToken(ctx context.Context, req *model.GenerateRefreshTokenReq) (resp *model.GenerateRefreshTokenResp, err error) {
	publicKeyCache := "jwk::public-key:*"

	if remaining := us.cache.RemainingTime(ctx, publicKeyCache); remaining <= 0 {
		err := us.generateJWKKey(ctx)
		if err != nil {
			return nil, err
		}
	}

	signer, ok := keyMap[config.GlobalCfg.JWK.KeyID]
	if !ok {
		err := us.generateJWKKey(ctx)
		if err != nil {
			return nil, err
		}
		signer = keyMap[config.GlobalCfg.JWK.KeyID]
	}

	claims := make(jwt.MapClaims)
	claims["sub"] = req.UserUID
	claims["iat"] = time.Now().UTC().Unix()
	claims["exp"] = time.Now().UTC().Add(time.Duration(req.ExpiredAt) * time.Second).Unix()

	data, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	obj, err := signer.Sign(data)
	if err != nil {
		return nil, err
	}

	tokenString, err := obj.CompactSerialize()
	if err != nil {
		return nil, err
	}

	return &model.GenerateRefreshTokenResp{
		Token: tokenString,
	}, nil
}

func (us *UsersServiceImpl) generateJWKKey(ctx context.Context) error {

	key, err := rsa.GenerateKey(rand.Reader, config.GlobalCfg.JWK.Size)
	if err != nil {
		return err
	}

	privateKeyID := config.GlobalCfg.JWK.KeyID

	privateKey := &jose.JSONWebKey{
		Key:       key,
		KeyID:     privateKeyID,
		Algorithm: config.GlobalCfg.JWK.Algorithm,
		Use:       config.GlobalCfg.JWK.Use,
	}

	thumb, err := privateKey.Thumbprint(crypto.SHA256)
	if err != nil {
		return err
	}

	publicKeyID := base64.URLEncoding.EncodeToString(thumb)
	privateKey.KeyID = publicKeyID

	publicKey := jose.JSONWebKey{
		Key:       key.Public(),
		KeyID:     publicKeyID,
		Algorithm: config.GlobalCfg.JWK.Algorithm,
		Use:       config.GlobalCfg.JWK.Use,
	}

	b, err := json.Marshal(publicKey)
	if err != nil {
		return err
	}

	publicKeyCache := fmt.Sprintf("jwk::public-key:%s", publicKeyID)
	expired := time.Now().Add(time.Duration(config.GlobalCfg.JWK.Expired))

	go us.cache.Set(ctx, publicKeyCache, string(b), int(expired.Unix()))

	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.SignatureAlgorithm(privateKey.Algorithm), Key: privateKey}, nil)
	if err != nil {
		return err
	}

	keyMap = make(map[string]jose.Signer)
	keyMap[config.GlobalCfg.JWK.KeyID] = signer

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

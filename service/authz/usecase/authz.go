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
	"github/yogabagas/join-app/service/authz/presenter"
	authzRepo "github/yogabagas/join-app/service/authz/repository"
	jwkRepo "github/yogabagas/join-app/service/jwk/repository"
	usersRepo "github/yogabagas/join-app/service/users/repository"
	"github/yogabagas/join-app/shared/constant"
	"github/yogabagas/join-app/shared/util"
	"log"
	"strings"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/golang-jwt/jwt"
)

var (
	keyMap map[string]jose.Signer
)

type AuthzServiceImpl struct {
	authzRepo authzRepo.AuthzRepository
	jwkRepo   jwkRepo.JWKRepository
	usersRepo usersRepo.UsersRepository
	cache     cache.Cache
	presenter presenter.AuthzPresenter
}

type AuthzService interface {
	Login(ctx context.Context, req service.LoginReq) (resp service.LoginResp, err error)
	Logout(ctx context.Context, req service.LogoutReq) error
	VerifyJWT(ctx context.Context, req service.VerifyTokenReq) (resp service.VerifyTokenResp, err error)
}

func NewAuthzService(repository sql.RepositoryRegistry, cache cache.Cache, presenter presenter.AuthzPresenter) AuthzService {
	return &AuthzServiceImpl{
		authzRepo: repository.AuthzRepository(),
		jwkRepo:   repository.JWKRepository(),
		usersRepo: repository.UserRepository(),
		cache:     cache,
		presenter: presenter,
	}
}

func (as *AuthzServiceImpl) Login(ctx context.Context, req service.LoginReq) (resp service.LoginResp, err error) {

	pwd, err := util.Hash(config.GlobalCfg.PasswordAlg, req.Password)
	if err != nil {
		return resp, err
	}

	user, err := as.usersRepo.ReadUserByEmailPassword(ctx, &model.ReadUserByEmailPasswordReq{
		Email:    req.Email,
		Password: util.Base64(pwd),
		RoleID:   constant.Role(req.RoleID).Int(),
	})
	if err != nil {
		return resp, err
	}

	key, err := as.jwkRepo.ReadUnexpiredKeyByID(ctx, &model.ReadUnexpiredKeyByIDReq{
		KeyID: user.RoleName,
	})
	if err != nil {
		return resp, err
	}

	accessToken, err := as.generateAndSignAccessToken(ctx, &model.GenerateAccessTokenReq{
		KeyID:      user.RoleName,
		UserUID:    user.UserUID,
		RoleUID:    user.RoleUID,
		IsValid:    key == nil,
		LastActive: user.LastActive.UTC().Unix(),
		ExpiredAt:  config.GlobalCfg.TokenExpiration,
	})
	if err != nil {
		return resp, err
	}

	refreshToken, err := as.generateAndSignRefreshToken(ctx, &model.GenerateRefreshTokenReq{
		KeyID:     user.RoleName,
		UserUID:   user.UserUID,
		ExpiredAt: (config.GlobalCfg.TokenExpiration + config.GlobalCfg.RefreshTokenExpiration),
	})
	if err != nil {
		return resp, err
	}

	cacheKeyLogin := fmt.Sprintf("auth::user-uid:%s", user.UserUID)
	tokenExp := time.Now().UTC().Add(time.Duration(config.GlobalCfg.TokenExpiration) * time.Second).Unix()

	err = as.cache.Set(ctx, cacheKeyLogin, true, int(tokenExp))
	if err != nil {
		log.Fatalln("error set cache auth", err)
	}

	return service.LoginResp{
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
	}, nil
}

func (as *AuthzServiceImpl) Logout(ctx context.Context, req service.LogoutReq) error {

	key := fmt.Sprintf("auth::user-uid:%s", req.UserUID)

	return as.cache.Delete(ctx, key)
}

func (as *AuthzServiceImpl) VerifyJWT(ctx context.Context, req service.VerifyTokenReq) (resp service.VerifyTokenResp, err error) {

	token, err := util.SplitBearer(req.Token)
	if err != nil {
		return resp, err
	}

	b, err := base64.RawStdEncoding.DecodeString(strings.Split(token, ".")[0])
	if err != nil {
		return resp, err
	}

	headerToken := make(map[string]string)
	if err := json.Unmarshal(b, &headerToken); err != nil {
		return resp, err
	}

	kid, ok := headerToken["kid"]
	if !ok {
		return resp, fmt.Errorf("token key ID is missing %s", kid)
	}

	object, err := jose.ParseSigned(token)
	if err != nil {
		return resp, err
	}

	keys, err := as.jwkRepo.ReadUnexpiredKeys(ctx)
	if err != nil {
		return
	}

	if len(keys) <= 0 {
		return resp, errors.New("key has expired")
	}

	var key interface{}
	for _, k := range keys {

		m := jose.JSONWebKey{}

		if err = json.Unmarshal(k.Key.([]byte), &m); err != nil {
			return
		}

		if m.KeyID == kid {
			key = m
		}

	}

	if key == nil {
		return resp, errors.New("key not found")
	}

	pb, err := object.Verify(key)
	if err != nil {
		log.Println("error verify object", err)
		return resp, err
	}

	payload := make(map[string]interface{})
	if err = json.Unmarshal(pb, &payload); err != nil {
		return resp, err
	}

	return as.presenter.VerifyJWT(ctx, payload)

}

func (as *AuthzServiceImpl) generateAndSignAccessToken(ctx context.Context, req *model.GenerateAccessTokenReq) (resp *model.GenerateAccessTokenResp, err error) {

	signer, ok := keyMap[req.KeyID]
	if !ok || !req.IsValid {
		err := as.generateJWKKey(ctx, req.KeyID)
		if err != nil {
			return nil, err
		}
		signer = keyMap[req.KeyID]
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

func (as *AuthzServiceImpl) generateAndSignRefreshToken(ctx context.Context, req *model.GenerateRefreshTokenReq) (resp *model.GenerateRefreshTokenResp, err error) {

	signer := keyMap[req.KeyID]

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

func (as *AuthzServiceImpl) generateJWKKey(ctx context.Context, keyID string) error {

	key, err := rsa.GenerateKey(rand.Reader, config.GlobalCfg.JWK.Size)
	if err != nil {
		return err
	}

	privateKeyID := keyID

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

	expired := time.Now().Add(time.Duration(config.GlobalCfg.JWK.Expired) * time.Hour).UTC()

	reqJWT := &model.JWK{
		ID:         privateKeyID,
		Key:        string(b),
		PrivateKey: privateKey,
		ExpiredAt:  expired,
	}

	err = as.jwkRepo.UpsertJWK(ctx, reqJWT)
	if err != nil {
		return err
	}

	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.SignatureAlgorithm(privateKey.Algorithm), Key: privateKey}, nil)
	if err != nil {
		return err
	}

	keyMap = make(map[string]jose.Signer)
	keyMap[keyID] = signer

	return nil
}

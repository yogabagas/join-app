package usecase

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"github/yogabagas/join-app/config"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/domain/repository/cache"
	"github/yogabagas/join-app/domain/repository/sql"
	"github/yogabagas/join-app/domain/service"
	authzRepo "github/yogabagas/join-app/service/authz/repository"
	jwkRepo "github/yogabagas/join-app/service/jwk/repository"
	usersRepo "github/yogabagas/join-app/service/users/repository"
	"github/yogabagas/join-app/shared/constant"
	"github/yogabagas/join-app/shared/util"
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
}

type AuthzService interface {
	Login(ctx context.Context, req service.LoginReq) (resp service.LoginResp, err error)
}

func NewAuthzService(repository sql.RepositoryRegistry, cache cache.Cache) AuthzService {
	return &AuthzServiceImpl{
		authzRepo: repository.AuthzRepository(),
		jwkRepo:   repository.JWKRepository(),
		usersRepo: repository.UserRepository(),
		cache:     cache,
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

	key, err := as.jwkRepo.ReadUnexpiredKey(ctx, &model.ReadUnexpiredKeyReq{
		Time: time.Now().UTC(),
	})
	if err != nil {
		return resp, err
	}

	accessToken, err := as.generateAndSignAccessToken(ctx, &model.GenerateAccessTokenReq{
		UserUID:    user.UserUID,
		RoleUID:    user.RoleUID,
		ExpiredAt:  config.GlobalCfg.TokenExpiration,
		LastActive: int(user.LastActive.UTC().Unix()),
	})
	if err != nil {
		return resp, err
	}

	refreshToken, err := as.generateAndSignRefreshToken(ctx, &model.GenerateRefreshTokenReq{
		UserUID:   user.UserUID,
		ExpiredAt: (config.GlobalCfg.TokenExpiration + config.GlobalCfg.RefreshTokenExpiration),
	})
	if err != nil {
		return resp, err
	}

	return service.LoginResp{
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
	}, nil
}

func (as *AuthzServiceImpl) generateAndSignAccessToken(ctx context.Context, req *model.GenerateAccessTokenReq) (resp *model.GenerateAccessTokenResp, err error) {

	signer, ok := keyMap[config.GlobalCfg.JWK.KeyID]
	if !ok {
		err := as.generateJWKKey(ctx)
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

func (as *AuthzServiceImpl) generateAndSignRefreshToken(ctx context.Context, req *model.GenerateRefreshTokenReq) (resp *model.GenerateRefreshTokenResp, err error) {

	signer, ok := keyMap[config.GlobalCfg.JWK.KeyID]
	if !ok {
		err := as.generateJWKKey(ctx)
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

func (as *AuthzServiceImpl) generateJWKKey(ctx context.Context) error {

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

	expired := time.Now().Add(time.Duration(config.GlobalCfg.JWK.Expired) * time.Second).UTC()

	reqJWT := &model.JWK{
		ID:        privateKeyID,
		Key:       string(b),
		ExpiredAt: expired,
	}

	err = as.jwkRepo.CreateJWK(ctx, reqJWT)
	if err != nil {
		return err
	}

	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.SignatureAlgorithm(privateKey.Algorithm), Key: privateKey}, nil)
	if err != nil {
		return err
	}

	keyMap = make(map[string]jose.Signer)
	keyMap[config.GlobalCfg.JWK.KeyID] = signer

	return nil
}

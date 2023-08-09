package service

import "time"

type JWTClaims struct {
	RoleUID string `json:"role_uid"`
	UserUID string `json:"user_uid"`
	Token   string `json:"token"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleID   int    `json:"role_id" validate:"required"`
}

type LoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type VerifyTokenReq struct {
	Token string
}

type VerifyTokenResp struct {
	Valid      bool      `json:"valid"`
	UserUID    string    `json:"user_uid"`
	RoleUID    string    `json:"role_uid"`
	LastActive time.Time `json:"last_active"`
	ExpiredAt  time.Time `json:"expired_at"`
}

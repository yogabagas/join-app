package service

type JWTClaims struct {
	RoleUID string `json:"role_uid"`
	UserUID string `json:"user_uid"`
	Token   string `json:"token"`
}

type AuthReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResp struct {
	Token        string
	RefreshToken string
}

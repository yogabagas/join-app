package service

type AuthReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResp struct {
	Token        string
	RefreshToken string
}

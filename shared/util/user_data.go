package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type UserData struct {
	RoleUUID string `json:"role_uuid"`
	UserUUID string `json:"user_id"`
	Token    string `json:"token"`
}

func (u UserData) GetUserData(r *http.Request) *UserData {
	tokenClaims := r.Context().Value("user_data").(*jwt.Token)
	claims, ok := tokenClaims.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Invalid claims type")
	}
	claimData := make(map[string]interface{})
	for key, value := range claims {
		claimData[key] = value
	}

	u.RoleUUID = claimData["role_uuid"].(string)
	u.UserUUID = claimData["user_uuid"].(string)

	return &u
}

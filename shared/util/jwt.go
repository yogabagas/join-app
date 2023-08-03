package util

import (
	"errors"
	"fmt"
	"github/yogabagas/join-app/domain/service"
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/dgrijalva/jwt-go"
)

func GetUserData(token string) (resp *service.JWTClaims, err error) {

	jwtToken, err := SplitBearer(token)
	if err != nil {
		return nil, err
	}

	tokenParse, err := new(jwt.Parser).ParseWithClaims(jwtToken, jwt.MapClaims{}, nil)
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenParse.Claims.(jwt.MapClaims); ok {
		if subject, ok := claims["sub"]; ok {
			resp.UserUID = fmt.Sprintf("%v", subject)
		}
		if role, ok := claims["role_uid"]; ok {
			resp.RoleUID = role.(string)
		}
	}
	return
}

func SplitBearer(token string) (string, error) {
	err := validation.Validate(token,
		validation.Required,
		validation.Match(regexp.MustCompile(`^(s|bearer|Bearer).([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_\-\+\/=]*)`)))
	if err != nil {
		return "", errors.New("unknown jwt format")
	}

	var bearer string

	switch {
	case strings.Contains(token, "bearer"):
		bearer = "bearer "
	case strings.Contains(token, "Bearer"):
		bearer = "Bearer "
	}

	token = strings.ReplaceAll(token, bearer, "")

	if token == "" {
		return "", errors.New("token is empty")
	}

	return token, nil

}

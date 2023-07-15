package middlewares

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github/yogabagas/print-in/config"
	"github/yogabagas/print-in/transport/rest/handler/response"
	"net/http"
	"strings"
)

var res = response.NewJSONResponse()

// AuthenticationMiddleware validates the JWT token.
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("authorization") != "" {
			if parseJwt(r.Header.Get("authorization"), w) {
				next.ServeHTTP(w, r)
			} else {
				res.SetError(response.ErrUnauthorized).SetMessage(errors.New("invalid authorized token").Error()).Send(w)
				return
			}
		} else {
			res.SetError(response.ErrUnauthorized).SetMessage(errors.New("an authorization header is required").Error()).Send(w)
			return
		}
	})
}

func parseJwt(authorizationHeader string, w http.ResponseWriter) (valid bool) {
	bearerToken := strings.Split(authorizationHeader, " ")
	if len(bearerToken) == 2 {
		token, _ := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(config.GlobalCfg.App.JwtSecret), nil
		})

		return token.Valid
	}
	return false
}

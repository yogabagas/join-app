package middlewares

import (
	"errors"
	"fmt"
	"github/yogabagas/join-app/config"
	"github/yogabagas/join-app/transport/rest/handler/response"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type MiddlewareImpl struct{}

type Middleware interface {
	AuthenticationMiddleware(next http.Handler) http.Handler
	CORSHandle(next http.Handler) http.Handler
}

func NewMiddleware() Middleware {
	return &MiddlewareImpl{}
}

// AuthenticationMiddleware validates the JWT token.
func (mi *MiddlewareImpl) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !mi.isWhitelist(r.URL.Path, r.Method) {
			res := response.NewJSONResponse()

			token := r.Header.Get("Authorization")

			if token == "" {
				res.SetError(response.ErrUnauthorized).SetMessage(errors.New("authorization header is required").Error()).Send(w)
				return
			}

			if !mi.parseJwt(token) {
				res.SetError(response.ErrUnauthorized).SetMessage(errors.New("invalid authorized token").Error()).Send(w)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (mi *MiddlewareImpl) isWhitelist(endpoint, method string) bool {
	mapAPI := make(map[string][]string)

	for _, v := range config.GlobalCfg.Whitelist.APIs {
		mapAPI[v.Endpoint] = append(mapAPI[v.Endpoint], v.Methods...)
	}

	if methods, ok := mapAPI[endpoint]; ok {
		for _, m := range methods {
			if method == m {
				return true
			}
		}
	}

	return false
}

func (mi *MiddlewareImpl) parseJwt(authorizationHeader string) (valid bool) {
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

package middlewares

import (
	"context"
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
		ctx := context.Background()
		if !mi.isWhitelist(r.URL.Path, r.Method) {
			res := response.NewJSONResponse()

			token := r.Header.Get("Authorization")

			if token == "" {
				res.SetError(response.ErrUnauthorized).SetMessage(errors.New("authorization header is required").Error()).Send(w)
				return
			}

			valid, newCtx := mi.parseJwt(token, r)
			if !valid {
				res.SetError(response.ErrUnauthorized).SetMessage(errors.New("invalid authorized token").Error()).Send(w)
				return
			}
			ctx = newCtx
		}

		next.ServeHTTP(w, r.WithContext(ctx))
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

func (mi *MiddlewareImpl) parseJwt(authorizationHeader string, r *http.Request) (valid bool, ctx context.Context) {
	bearerToken := strings.Split(authorizationHeader, " ")
	if len(bearerToken) == 2 {
		token, _ := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(config.GlobalCfg.App.JWTSecret), nil
		})

		ctx = context.WithValue(r.Context(), "user_data", token)

		return token.Valid, ctx
	}
	return false, ctx
}

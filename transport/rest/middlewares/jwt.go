package middlewares

import (
	"errors"
	"fmt"
	"github/yogabagas/print-in/config"
	"github/yogabagas/print-in/transport/rest/handler/response"
	"net/http"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type MiddlewareImpl struct{}

type Middleware interface {
	AuthenticationMiddleware(next http.Handler) http.Handler
}

func NewMiddleware() Middleware {
	return &MiddlewareImpl{}
}

var res = response.NewJSONResponse()

// AuthenticationMiddleware validates the JWT token.
func (mi *MiddlewareImpl) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		isRegister := r.URL.Path == "/v1/users" && r.Method == http.MethodPost

		if token == "" && !isRegister {
			res.SetError(response.ErrUnauthorized).SetMessage(errors.New("An Authorization Header is required").Error()).Send(w)
			return
		}

		if !mi.parseJwt(token) && !isRegister {
			res.SetError(response.ErrUnauthorized).SetMessage(errors.New("Invalid Authorized Token").Error()).Send(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (mi *MiddlewareImpl) parseJwt(authorizationHeader string) (valid bool) {
	bearerToken := strings.Split(authorizationHeader, " ")
	if len(bearerToken) == 2 {
		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("token is not compatible")
			}
			return []byte(config.GlobalCfg.App.JwtSecret), nil
		})
		if err != nil {
			return false
		}

		return token.Valid
	}
	return false
}

func DecodeToken(token string) (claim service.ClaimsJWT, err error) {

	jwtToken, err := SplitBearer(token)
	if err != nil {
		return claim, err
	}

	tokenParse, _, err := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})
	if err != nil {
		return claim, errors.New(err.Error())
	}

	if claims, ok := tokenParse.Claims.(jwt.MapClaims); ok {
		if user, ok := claims["userId"]; ok {
			claim.UserID = fmt.Sprintf("%v", user)
		}
		if policy, ok := claims["policy"]; ok {
			claim.Policy = policy.(string)
		}
		if uName, ok := claims["username"]; ok {
			claim.UserName = uName.(string)
		}
		if eMail, ok := claims["email"]; ok {
			claim.Email = eMail.(string)
		}
		if clientID, ok := claims["clientId"]; ok {
			claim.ClientID = clientID.(string)
		}
		if typeInternal, ok := claims["type"]; ok {
			claim.Type = typeInternal.(string)
		}
		return
	}
	return claim, errors.New("map claims token is empty")
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

package v1

import (
	"github/yogabagas/print-in/transport/rest/handler"
	"github/yogabagas/print-in/transport/rest/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func NewUsersV1(h handler.HandlerImpl, r *mux.Router) {
	r.Use(middlewares.AuthenticationMiddleware)
	r.HandleFunc("/users", h.CreateUsers).Methods(http.MethodPost)

}

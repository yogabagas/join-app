package v1

import (
	"github/yogabagas/print-in/transport/rest/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewUsersV1(h handler.HandlerImpl, r *mux.Router) {
	r.HandleFunc("/users", h.CreateUsers).Methods(http.MethodPost)
}

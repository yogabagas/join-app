package v1

import (
	"github/yogabagas/print-in/transport/rest/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRolesV1(h handler.HandlerImpl, r *mux.Router) {
	r.HandleFunc("/roles", h.CreateRoles).Methods(http.MethodPost)
}

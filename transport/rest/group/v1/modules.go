package v1

import (
	"github/yogabagas/join-app/transport/rest/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewModulesV1(h handler.HandlerImpl, r *mux.Router) {
	r.HandleFunc("/module/create", h.CreateModules).Methods(http.MethodPost)
	r.HandleFunc("/modules", h.GetCourses).Methods(http.MethodGet)
	r.HandleFunc("/modules/:id", h.UpdateCourses).Methods(http.MethodPut)
	r.HandleFunc("/modules/:id", h.CoursesByID).Methods(http.MethodGet)
	r.HandleFunc("/modules/:id", h.DeleteCourse).Methods(http.MethodDelete)
}

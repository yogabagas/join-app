package v1

import (
	"github/yogabagas/join-app/transport/rest/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewCoursesV1(h handler.HandlerImpl, r *mux.Router) {
	r.HandleFunc("/courses", h.CreateCourses).Methods(http.MethodPost)
	r.HandleFunc("/courses", h.GetCourses).Methods(http.MethodGet)
	r.HandleFunc("/courses/:id", h.UpdateCourses).Methods(http.MethodPut)
	r.HandleFunc("/courses/:id", h.CoursesByID).Methods(http.MethodGet)
	r.HandleFunc("/courses/:id", h.DeleteCourse).Methods(http.MethodDelete)
	r.HandleFunc("/:courses_detail_id/:courses_id", h.GetCourseDetail).Methods(http.MethodGet)
}

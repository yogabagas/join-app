package handler

import (
	"encoding/json"
	"github/yogabagas/join-app/domain/service"
	"github/yogabagas/join-app/transport/rest/handler/response"
	"net/http"
)

// CreateCourses handler
// @Summary Create New Courses
// @Description New Courses Registration
// @Tags Courses
// @Produce json
// @Param users body service.CreateCoursesReq true "Request Create Resources"
// @Success 200 {object} response.JSONResponse().APIStatusCreated()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/courses [POST]
func (h *HandlerImpl) CreateCourses(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.CreateCoursesReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	if err := h.Controller.CoursesController.CreateCourses(r.Context(), req); err != nil {
		res.SetError(response.ErrInternalServerError).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusCreated().Send(w)

}

// GetCourses handler
// @Summary Get All Courses
// @Description New Resources Registration
// @Tags Resources
// @Produce json
// @Param users body service.CreateResourcesReq true "Request Create Resources"
// @Success 200 {object} response.JSONResponse().APIStatusCreated()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/courses [GET]
func (h *HandlerImpl) GetCourses(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.CreateResourcesReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	err := h.Controller.ResourcesController.CreateResources(r.Context(), req)
	if err != nil {
		res.SetError(response.ErrInternalServerError).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusCreated().Send(w)

}

// GetCourses handler
// @Summary Get All Courses
// @Description New Resources Registration
// @Tags Resources
// @Produce json
// @Param users body service.CreateResourcesReq true "Request Create Resources"
// @Success 200 {object} response.JSONResponse().APIStatusCreated()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/courses [GET]
func (h *HandlerImpl) UpdateCourses(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.CreateResourcesReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	err := h.Controller.ResourcesController.CreateResources(r.Context(), req)
	if err != nil {
		res.SetError(response.ErrInternalServerError).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusCreated().Send(w)

}

// GetCourses handler
// @Summary Get All Courses
// @Description New Resources Registration
// @Tags Resources
// @Produce json
// @Param users body service.CreateResourcesReq true "Request Create Resources"
// @Success 200 {object} response.JSONResponse().APIStatusCreated()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/courses [GET]
func (h *HandlerImpl) CoursesByID(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.CreateResourcesReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	err := h.Controller.ResourcesController.CreateResources(r.Context(), req)
	if err != nil {
		res.SetError(response.ErrInternalServerError).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusCreated().Send(w)

}

// GetCourses handler
// @Summary Get All Courses
// @Description New Resources Registration
// @Tags Resources
// @Produce json
// @Param users body service.CreateResourcesReq true "Request Create Resources"
// @Success 200 {object} response.JSONResponse().APIStatusCreated()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/courses [GET]
func (h *HandlerImpl) DeleteCourse(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.CreateResourcesReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	err := h.Controller.ResourcesController.CreateResources(r.Context(), req)
	if err != nil {
		res.SetError(response.ErrInternalServerError).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusCreated().Send(w)

}

// GetCourseDetail handler
// @Summary Get All Courses
// @Description New Resources Registration
// @Tags Resources
// @Produce json
// @Param users body service.CreateResourcesReq true "Request Create Resources"
// @Success 200 {object} response.JSONResponse().APIStatusCreated()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/courses [GET]
func (h *HandlerImpl) GetCourseDetail(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.CreateResourcesReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	err := h.Controller.ResourcesController.CreateResources(r.Context(), req)
	if err != nil {
		res.SetError(response.ErrInternalServerError).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusCreated().Send(w)

}

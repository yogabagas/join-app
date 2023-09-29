package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github/yogabagas/join-app/domain/service"
	"github/yogabagas/join-app/shared/util"
	"github/yogabagas/join-app/transport/rest/handler/response"
	"net/http"
	"strconv"
)

// CreateModules handler
// @Summary Create New Modules
// @Description New Modules Registration
// @Tags Modules
// @Produce json
// @Param users body service.CreateModulesReq true "Request Create Modules"
// @Success 200 {object} response.JSONResponse().APIStatusCreated()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/modules/create [POST]
func (h *HandlerImpl) CreateModules(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.CreateModulesReq

	req, err := req.SetCreateModuleReq(r)
	if err != nil {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	userData := new(util.UserData)
	userData = userData.GetUserData(r)

	if err := h.Controller.ModulesController.CreateModules(r.Context(), req, userData); err != nil {
		res.SetError(response.ErrInternalServerError).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusCreated().Send(w)

}

// GetModules handler
// @Summary Get All Modules
// @Description New Modules Registration
// @Tags Modules
// @Produce json
// @Param users body service.CreateModulesReq true "Request Create Modules"
// @Success 200 {object} response.JSONResponse().APIStatusCreated()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/modules [GET]
func (h *HandlerImpl) GetModulesWithPagination(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodGet {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.GetModulesWithPaginationReq

	if name := r.URL.Query().Get("name"); name != "" {
		req.Name = name
	}

	var limitToInt int
	if limit := r.URL.Query().Get("limit"); limit != "" {
		limitToInt, _ = strconv.Atoi(limit)
	}

	if limitToInt <= 0 {
		limitToInt = 10
	}
	req.Limit = limitToInt

	var pageToInt int
	if page := r.URL.Query().Get("page"); page != "" {
		pageToInt, _ = strconv.Atoi(page)
	}

	if pageToInt <= 0 {
		pageToInt = 1
	}
	req.Page = pageToInt

	resp, err := h.Controller.ModulesController.GetModulesWithPagination(r.Context(), req)
	if err != nil {
		res.SetError(response.ErrInternalServerError).SetMessage(err.Error()).Send(w)
		return
	}

	res.SetData(resp).Send(w)

}

// GetModules handler
// @Summary Get All Courses
// @Description New Resources Registration
// @Tags Resources
// @Produce json
// @Param users body service.CreateResourcesReq true "Request Create Resources"
// @Success 200 {object} response.JSONResponse().APIStatusCreated()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/modules [GET]
func (h *HandlerImpl) UpdateCourses(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.CreateModulesReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	userData := new(util.UserData)
	userData = userData.GetUserData(r)

	err := h.Controller.ModulesController.UpdateModules(r.Context(), req, userData)
	if err != nil {
		res.SetError(response.ErrInternalServerError).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusCreated().Send(w)

}

// GetModules handler
// @Summary Get All Courses
// @Description New Resources Registration
// @Tags Resources
// @Produce json
// @Param users body service.CreateResourcesReq true "Request Create Resources"
// @Success 200 {object} response.JSONResponse().APIStatusCreated()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/modules [GET]
func (h *HandlerImpl) DeleteCourse(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodDelete {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}
	vars := mux.Vars(r)
	userData := new(util.UserData)
	userData = userData.GetUserData(r)

	err := h.Controller.ModulesController.DeleteModules(r.Context(), vars["uid"], *userData)
	if err != nil {
		res.SetError(response.ErrInternalServerError).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusCreated().Send(w)

}

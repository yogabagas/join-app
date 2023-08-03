package handler

import (
	"encoding/json"
	"github/yogabagas/join-app/domain/service"
	"github/yogabagas/join-app/shared/util"
	"github/yogabagas/join-app/transport/rest/handler/response"
	"net/http"
	"strconv"
)

// CreateUsers handler
// @Summary Create New User
// @Description New User Registration
// @Tags Users
// @Produce json
// @Param users body service.CreateUsersReq true "Request Create User"
// @Success 200 {object} response.JSONResponse().APIStatusCreated()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/users [POST]
func (h *HandlerImpl) CreateUsers(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.CreateUsersReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	err := h.Controller.UsersController.CreateUsers(r.Context(), req)
	if err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusCreated().Send(w)
}

// Login handler
// @Summary Login
// @Description Login registration endpoint
// @Tags Users V1.0
// @Produce json
// @Param users body service.LoginReq true "Request Login"
// @Success 200 {object} response.JSONResponse().APIStatusSuccess()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/login [POST]
func (h *HandlerImpl) Login(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.LoginReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	if !util.ValidateEmail(req.Email) {
		res.SetError(response.ErrBadRequest).SetMessage("Invalid Email Format").Send(w)
		return
	}

	user, err := h.Controller.UsersController.Login(r.Context(), req)
	if err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusSuccess().SetResult(user).Send(w)
}

// Logout handler
// @Summary Login
// @Description Login registration endpoint
// @Tags Users V1.0
// @Produce json
// @Param users body true "Request Logout"
// @Success 200 {object} response.JSONResponse().APIStatusSuccess()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/logout [POST]
func (h *HandlerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	res := response.NewJSONResponse()

	if r.Method != http.MethodDelete {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	userData := new(util.UserData)
	userData = userData.GetUserData(r)

	_, err := h.Controller.UsersController.Logout(r.Context(), userData.UserUUID)

	if err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	res.SetStatusCode(204).Send(w)
}

// GetUsersWithPagination handler
// @Summary GetUsersWithPagination
// @Description GetUsersWithPagination for get users detail with limit
// @Tags Users
// @Produce json
// @Param name query string false "user fullname e.g John Doe"
// @Param limit query int false "limit data; default 10"
// @Param page query int false "number of page; default 1"
// @Success 200 {object} response.JSONResponse{data=service.GetUsersWithPaginationResp}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/users [GET]
func (h *HandlerImpl) GetUsersWithPagination(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodGet {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.GetUsersWithPaginationReq

	if name := r.URL.Query().Get("name"); name != "" {
		req.Fullname = name
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

	resp, err := h.Controller.UsersController.GetUsersWithPagination(r.Context(), req)
	if err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	res.SetData(resp).Send(w)
}

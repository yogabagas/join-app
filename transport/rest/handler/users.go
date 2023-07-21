package handler

import (
	"encoding/json"
	"github/yogabagas/print-in/domain/service"
	"github/yogabagas/print-in/shared/util"
	"github/yogabagas/print-in/transport/rest/handler/response"
	"net/http"
)

// CreateUsers handler
// @Summary Create New User
// @Description User registration endpoint
// @Tags Users V1.0
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
// @Param users body service.CreateUsersReq true "Request Create User"
// @Success 200 {object} response.JSONResponse().APIStatusCreated()
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
	//var req service.CreateUsersReq

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

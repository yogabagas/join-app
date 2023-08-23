package handler

import (
	"encoding/json"
	"github/yogabagas/join-app/domain/service"
	"github/yogabagas/join-app/shared/util"
	"github/yogabagas/join-app/transport/rest/handler/response"
	"net/http"
)

// Login handler
// @Summary Login
// @Description Login endpoint
// @Tags Users
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

	user, err := h.Controller.AuthzController.Login(r.Context(), req)
	if err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusSuccess().SetResult(user).Send(w)
}

// Logout handler
// @Summary Lgout
// @Description Logout endpoint
// @Tags Users
// @Produce json
// @Security ApiKeyAuth
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

	token := r.Header.Get("Authorization")

	claims, err := h.Controller.AuthzController.VerifyJWT(r.Context(), service.VerifyTokenReq{
		Token: token,
	})
	if err != nil {
		res.SetError(response.ErrUnauthorized).SetMessage(err.Error()).Send(w)
		return
	}

	req := service.LogoutReq{
		UserUID: claims.UserUID,
	}

	err = h.Controller.AuthzController.Logout(r.Context(), req)

	if err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusNoContent().Send(w)
}

func (h *HandlerImpl) VerifyJWT(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodGet {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		res.SetError(response.ErrBadRequest).SetMessage("token can't be empty").Send(w)
		return
	}

	req := service.VerifyTokenReq{
		Token: token,
	}

	resp, err := h.Controller.AuthzController.VerifyJWT(r.Context(), req)
	if err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusAccepted().SetData(resp).Send(w)
}
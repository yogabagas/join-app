package handler

import (
	"encoding/json"
	"errors"
	"github/yogabagas/join-app/domain/service"
	"github/yogabagas/join-app/shared/constant"
	"github/yogabagas/join-app/shared/util"
	"github/yogabagas/join-app/transport/rest/handler/response"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *HandlerImpl) UpsertAccess(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPut {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.UpsertAccessReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	err := h.Controller.AccessController.UpsertAccess(r.Context(), req)
	if err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusCreated().Send(w)
}

func (h *HandlerImpl) GetAccessByRoleUID(w http.ResponseWriter, r *http.Request) {
	res := response.NewJSONResponse()

	if r.Method != http.MethodGet {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		res.SetError(response.ErrUnauthorized).SetMessage(errors.New("unauthorized access").Error()).Send(w)
		return
	}

	claims, err := util.GetUserData(token)
	if err != nil {
		res.SetError(response.ErrUnauthorized).SetMessage(errors.New("jwt is invalid format").Error()).Send(w)
		return
	}

	vars := mux.Vars(r)
	t, ok := vars["type"]
	if !ok {
		res.SetError(response.ErrBadRequest).SetMessage(errors.New("resources type is missing").Error()).Send(w)
		return
	}

	req := service.GetAccessByRoleUIDReq{
		RoleUID: claims.RoleUID,
		Type:    constant.ResourceTypeAtoi(t).Int(),
	}

	resp, err := h.Controller.AccessController.GetAccessByRoleUID(r.Context(), req)
	if err != nil {
		res.SetError(response.ErrInternalServerError).SetMessage(err.Error()).Send(w)
		return
	}

	res.SetData(resp).Send(w)
}

package handler

import (
	"encoding/json"
	"github/yogabagas/print-in/domain/service"
	"github/yogabagas/print-in/transport/rest/handler/response"
	"net/http"
)

// CreateRoles handler
// @Summary Create New Roles
// @Description Roles registration endpoint
// @Tags Roles V1.0
// @Produce json
// @Security ApiKeyAuth
// @Param users body service.CreateRolesReq true "Request Roles User"
// @Success 200 {object} response.JSONResponse().APIStatusCreated()
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /v1/roles [POST]
func (h *HandlerImpl) CreateRoles(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.CreateRolesReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(w)
		return
	}

	if err := h.Controller.RolesController.CreateRoles(r.Context(), req); err != nil {
		res.SetError(response.ErrInternalServerError).SetMessage(err.Error()).Send(w)
		return
	}

	res.APIStatusCreated().Send(w)
}

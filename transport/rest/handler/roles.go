package handler

import (
	"encoding/json"
	"github/yogabagas/print-in/domain/service"
	"github/yogabagas/print-in/transport/rest/handler/response"
	"net/http"
)

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

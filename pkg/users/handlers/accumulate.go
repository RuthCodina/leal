package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/leal/pkg/helpers"
	"github.com/leal/pkg/users/handlers/dtos"
)

func (h *Handler) AccumulatePoints(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	accDto := dtos.AccumulatePointsDTO{}
	err := json.NewDecoder(r.Body).Decode(&accDto)
	if err != nil {
		helpers.ErrInvalidJSON.Send(w)
		return
	}

	acc, err := h.UserService.AccumulatePoints(ctx, accDto.Name, accDto.Id, accDto.Sucursal, accDto.Payment)

	if err != nil {
		helpers.ErrBadRequest.Send(w)
		return
	}

	helpers.Success(&acc, http.StatusOK).Send(w)
}

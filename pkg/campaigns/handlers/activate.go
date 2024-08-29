package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/leal/pkg/campaigns/handlers/dtos"
)

func (h *Handler) Activate(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	activateDto := dtos.ActivateDTO{}
	err := json.NewDecoder(r.Body).Decode(&activateDto)
	if err != nil {
		ErrInvalidJSON.Send(w)
		return
	}

	err = h.CampaignService.Activate(ctx, activateDto.Name, activateDto.Id)

	if err != nil {
		ErrBadRequest.Send(w)
		return
	}

	Success(activateDto, http.StatusOK).Send(w)
}
func (h *Handler) Inactivate(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	activateDto := dtos.ActivateDTO{}
	err := json.NewDecoder(r.Body).Decode(&activateDto)
	if err != nil {
		ErrInvalidJSON.Send(w)
		return
	}

	err = h.CampaignService.Inactivate(ctx, activateDto.Name, activateDto.Id)

	if err != nil {
		ErrBadRequest.Send(w)
		return
	}

	Success(activateDto, http.StatusOK).Send(w)
}

package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/leal/pkg/campaigns/handlers/dtos"
	"github.com/leal/pkg/helpers"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	campaignDto := dtos.CampaignCreateDTO{}
	err := json.NewDecoder(r.Body).Decode(&campaignDto)
	if err != nil {
		helpers.ErrInvalidJSON.Send(w)
		return
	}

	campaign, err := h.CampaignService.Create(ctx, campaignDto.Name, campaignDto.Description, campaignDto.BranchOffice)

	if err != nil {
		helpers.ErrBadRequest.Send(w)
		return
	}

	log.Println(campaign)

	helpers.Success(&campaign, http.StatusOK).Send(w)

}

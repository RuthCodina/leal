package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/leal/pkg/campaigns/handlers/dtos"
)

func (h *Handler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	campaignDto := dtos.CampaignCreateDTO{}
	err := json.NewDecoder(r.Body).Decode(&campaignDto)
	if err != nil {
		ErrInvalidJSON.Send(w)
		return
	}

	campaign, err := h.CampaignService.Create(ctx, campaignDto.Name, campaignDto.Description, campaignDto.BranchOffice)

	if err != nil {
		ErrBadRequest.Send(w)
		return
	}

	log.Println(campaign)

	Success(&campaign, http.StatusOK).Send(w)
	/*
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(campaign)
	*/
}

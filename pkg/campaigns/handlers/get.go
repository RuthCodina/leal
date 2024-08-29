package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/leal/pkg/campaigns/handlers/dtos"
	"github.com/leal/pkg/helpers"
)

func (h *Handler) GetAllByCommerce(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	commerceDto := dtos.GetAllByCommerceDTO{}
	err := json.NewDecoder(r.Body).Decode(&commerceDto)
	if err != nil {
		helpers.ErrInvalidJSON.Send(w)
		return
	}

	campaigns, err := h.CampaignService.GetAllByCommerce(ctx, commerceDto.CommerceId, commerceDto.Activate)
	if err != nil {
		helpers.ErrBadRequest.Send(w)
		return
	}

	log.Println(campaigns)

	helpers.Success(&campaigns, http.StatusOK).Send(w)
}

func (h *Handler) GetAllByBranch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	branchDto := dtos.GetAllByBranchDTO{}
	err := json.NewDecoder(r.Body).Decode(&branchDto)
	if err != nil {
		helpers.ErrInvalidJSON.Send(w)
		return
	}

	campaigns, err := h.CampaignService.GetAllByBranch(ctx, branchDto.BranchId, branchDto.Activate)
	if err != nil {
		helpers.ErrBadRequest.Send(w)
		return
	}

	log.Println(campaigns)

	helpers.Success(&campaigns, http.StatusOK).Send(w)
}

package handlers

import (
	"github.com/leal/pkg/campaigns/domain"
)

type Handler struct {
	CampaignService domain.CampaignService
}

func NewCampaignHandler(s domain.CampaignService) *Handler {
	return &Handler{
		CampaignService: s,
	}
}

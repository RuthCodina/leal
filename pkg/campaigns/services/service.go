package services

import "github.com/leal/pkg/campaigns/domain"

var _ domain.CampaignService = &CampaignService{}

type CampaignService struct {
	campaignRepository domain.CampaignRepository
}

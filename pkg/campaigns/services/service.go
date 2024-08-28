package services

import "github.com/leal/pkg/campaigns/domain"

var _ domain.CampaignService = &Service{}

type Service struct {
	Repository domain.CampaignRepository
}

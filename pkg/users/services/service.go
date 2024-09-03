package services

import (
	cd "github.com/leal/pkg/campaigns/domain"
	"github.com/leal/pkg/users/domain"
)

var _ domain.UserService = &Service{}

type Service struct {
	Repository  domain.UserRepository
	CampaignSrv cd.CampaignService
}

func NewUserService(repo domain.UserRepository, srv cd.CampaignService) *Service {
	return &Service{
		Repository:  repo,
		CampaignSrv: srv,
	}
}

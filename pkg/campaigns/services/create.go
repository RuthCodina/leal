package services

import (
	"context"
	"errors"
	"log"

	"github.com/leal/pkg/campaigns/domain"
)

func (s *CampaignService) Create(ctx context.Context, name string) (*domain.Campaign, error) {
	campaign, err := s.campaignRepository.Create(ctx, name)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateKey) {
			log.Println("duplicate_key_error")
			srvErr := domain.Error{
				Code: domain.ErrCodeDuplicateKey,
				Msg:  "error creating campaign: duplicate key error",
			}
			return nil, srvErr
		}
		log.Println(err.Error())
		return nil, err
	}
	return campaign, nil
}

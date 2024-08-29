package services

import (
	"context"
	"errors"
	"log"

	"github.com/leal/pkg/campaigns/domain"
	"github.com/leal/pkg/helpers"
)

func (s *Service) Create(ctx context.Context, name, description string, branchOffice int) (*domain.Campaign, error) {
	campaign, err := s.Repository.Create(ctx, name, description, branchOffice)
	if err != nil {
		if errors.Is(err, helpers.ErrDuplicateKey) {
			log.Println("duplicate_key_error")
			srvErr := helpers.Error{
				Code: helpers.ErrCodeDuplicateKey,
				Msg:  "error creating campaign: duplicate key error",
			}
			return nil, srvErr
		}
		log.Println(err.Error())
		return nil, err
	}
	return campaign, nil
}

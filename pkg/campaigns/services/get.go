package services

import (
	"context"
	"errors"
	"log"

	"github.com/leal/pkg/campaigns/domain"
	"github.com/leal/pkg/helpers"
)

func (s *Service) GetAllByBranch(ctx context.Context, branchId int, active bool) (campaigns *[]domain.Campaign, err error) {
	campaigns, err = s.Repository.GetAllByBranch(ctx, branchId, active)
	if err != nil {
		if errors.Is(err, helpers.ErrNotFound) {
			log.Println("not_found_error")
			srvErr := helpers.Error{
				Code: helpers.ErrCodeNotFound,
				Msg:  "error getting campaigns by branch: not found error",
			}
			return campaigns, srvErr
		}
		log.Println(err.Error())
		return campaigns, err
	}

	return campaigns, nil
}
func (s *Service) GetAllByCommerce(ctx context.Context, commerceId int, active bool) (campaigns *[]domain.Campaign, err error) {
	campaigns, err = s.Repository.GetAllByCommerce(ctx, commerceId, active)
	if err != nil {
		if errors.Is(err, helpers.ErrNotFound) {
			log.Println("not_found_error")
			srvErr := helpers.Error{
				Code: helpers.ErrCodeNotFound,
				Msg:  "error getting campaigns by commerce: not found error",
			}
			return campaigns, srvErr
		}
		log.Println(err.Error())
		return campaigns, err
	}

	return campaigns, nil
}

func (s *Service) Status(ctx context.Context, branchId int) (*domain.CampaignStatus, error) {
	return s.Repository.Status(ctx, branchId)
}

func CampaignName(ctx context.Context, name string) interface{} {
	var campaign interface{}
	for i, v := range Names {
		if i == name {
			campaign = v.(func(ctx context.Context, args ...interface{}) interface{})
		}
	}
	return campaign
}

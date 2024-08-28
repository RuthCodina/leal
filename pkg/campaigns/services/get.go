package services

import (
	"context"
	"errors"
	"log"

	"github.com/leal/pkg/campaigns/domain"
)

func (s *Service) GetAllByBranch(ctx context.Context, name string, branchId int, active bool) (campaigns *[]domain.Campaign, err error) {
	campaigns, err = s.Repository.GetAllByBranch(ctx, name, branchId, active)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			log.Println("not_found_error")
			srvErr := domain.Error{
				Code: domain.ErrCodeNotFound,
				Msg:  "error getting campaigns: not found error",
			}
			return campaigns, srvErr
		}
		log.Println(err.Error())
		return campaigns, err
	}

	return campaigns, nil
}
func (s *Service) GetAllByCommerce(ctx context.Context, name string, commerceId int, active bool) (campaigns *[]domain.Campaign, err error) {
	campaigns, err = s.Repository.GetAllByCommerce(ctx, name, commerceId, active)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			log.Println("not_found_error")
			srvErr := domain.Error{
				Code: domain.ErrCodeNotFound,
				Msg:  "error getting campaigns: not found error",
			}
			return campaigns, srvErr
		}
		log.Println(err.Error())
		return campaigns, err
	}

	return campaigns, nil
}

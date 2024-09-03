package services

import (
	"context"
	"errors"
	"log"

	"github.com/leal/pkg/helpers"
	"github.com/leal/pkg/users/domain"
)

func (s *Service) AccumulatePoints(ctx context.Context, name string, userId int, payment domain.UserPayment) (acc *domain.UserAccumulated, err error) {
	campaignStatus, err := s.CampaignSrv.Status(ctx, payment.BranchId)
	if err != nil {
		return nil, err
	}

	promotion, err := getCampaignRate(ctx, payment.BranchId)
	if err != nil {
		return nil, err
	}

	if payment.PayDate.After(campaignStatus.ActiveDate) && payment.PayDate.Before(campaignStatus.InactiveDate) {
		acc, err = s.Repository.AccumulatePoints(ctx, name, userId, promotion)
		if err != nil {
			if errors.Is(err, helpers.ErrTimeout) {
				log.Println("time_out_error")
				srvErr := helpers.Error{
					Code: helpers.ErrCodeTimeout,
					Msg:  "error calculating cashback: time out error",
				}
				return nil, srvErr
			}
			log.Println(err.Error())
			return nil, err
		}
	}

	_, err = s.Repository.UpdateUser(ctx, userId, name, acc)
	if err != nil {
		if errors.Is(err, helpers.ErrTimeout) {
			log.Println("time_out_error")
			srvErr := helpers.Error{
				Code: helpers.ErrCodeTimeout,
				Msg:  "error upating user: time out Error",
			}
			return nil, srvErr
		}
		log.Println(err.Error())
		return nil, err
	}

	return acc, nil
}

func getCampaignRate(ctx context.Context, branchId int) (float32, error) {
	var s Service
	campaignName, err := s.Repository.GetCampaignName(ctx, branchId)
	if err != nil {
		return 0, err
	}
	promotion := s.CampaignSrv.CampaignRate(ctx, campaignName)
	return promotion, nil
}

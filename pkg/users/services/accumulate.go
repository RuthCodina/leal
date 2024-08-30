package services

import (
	"context"
	"errors"
	"log"

	"github.com/leal/pkg/helpers"
	"github.com/leal/pkg/users/domain"
	"github.com/shopspring/decimal"
)

func (s *Service) AccumulatePoints(ctx context.Context, name string, userId int, branchId int, payment domain.UserPayment) (*decimal.Decimal, error) {
	acc, err := s.Repository.AccumulatePoints(ctx, name, userId, branchId, payment)
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

	return acc, nil
}

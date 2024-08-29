package services

import (
	"context"
	"errors"
	"log"

	"github.com/leal/pkg/helpers"
	"github.com/shopspring/decimal"
)

func (s *Service) AccumulatePoints(ctx context.Context, userId int, branchId int, paid decimal.Decimal) (decimal.Decimal, error) {
	acc, err := s.Repository.AccumulatePoints(ctx, userId, branchId, paid)
	if err != nil {
		if errors.Is(err, helpers.ErrTimeout) {
			log.Println("time_out_error")
			srvErr := helpers.Error{
				Code: helpers.ErrCodeTimeout,
				Msg:  "error calculating cashback: time out error",
			}
			return decimal.NewFromInt(0), srvErr
		}
		log.Println(err.Error())
		return decimal.NewFromInt(0), err
	}

	return acc, nil
}

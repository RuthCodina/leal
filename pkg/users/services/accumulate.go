package services

import "github.com/shopspring/decimal"

func (s *Service) AccumulatePoints(userId int, branchId int) (decimal.Decimal, error) {
	return decimal.NewFromInt(0), nil
}

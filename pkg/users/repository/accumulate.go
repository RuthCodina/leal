package repository

import "github.com/shopspring/decimal"

func (r *Repository) AccumulatePoints(userId int, branchId int) (decimal.Decimal, error) {
	return decimal.NewFromInt(0), nil
}

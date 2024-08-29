package repository

import (
	"context"

	"github.com/shopspring/decimal"
)

func (r *Repository) AccumulatePoints(ctx context.Context, userId int, branchId int, paid decimal.Decimal) (decimal.Decimal, error) {
	return decimal.NewFromInt(0), nil
}

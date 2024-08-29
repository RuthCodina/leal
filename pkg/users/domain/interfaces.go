package domain

import (
	"context"

	"github.com/shopspring/decimal"
)

type UserRepository interface {
	AccumulatePoints(ctx context.Context, userId int, branchId int, paid decimal.Decimal) (decimal.Decimal, error)
}

type UserService interface {
	AccumulatePoints(ctx context.Context, userId int, branchId int, paid decimal.Decimal) (decimal.Decimal, error)
}

package domain

import (
	"context"

	"github.com/shopspring/decimal"
)

type UserRepository interface {
	AccumulatePoints(ctx context.Context, name string, userId int, branchId int, payment UserPayment) (*decimal.Decimal, error)
}

type UserService interface {
	AccumulatePoints(ctx context.Context, name string, userId int, branchId int, payment UserPayment) (*decimal.Decimal, error)
}

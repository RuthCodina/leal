package domain

import (
	"github.com/shopspring/decimal"
)

type UserRepository interface {
	AccumulatePoints(userId int, branchId int) (decimal.Decimal, error)
}

type UserService interface {
	AccumulatePoints(userId int, branchId int) (decimal.Decimal, error)
}

type UserHandler interface {
	AccumulatePoints(userId int, branchId int) (decimal.Decimal, error)
}

package dtos

import (
	"github.com/shopspring/decimal"
)

type AccumulatePointsDTO struct {
	Id      int        `json:"user_id", validate:"required"`
	Name    string     `json:"name", validate:"required"`
	Payment PaymentDTO `json:"payment", validate:"required"`
}

type PaymentDTO struct {
	Amount   decimal.Decimal `json:"amount", validate:"required"`
	PayDate  string          `json:"pay_date", validate:"required"`
	BranchId int             `json:"branch_id", validate:"required"`
}

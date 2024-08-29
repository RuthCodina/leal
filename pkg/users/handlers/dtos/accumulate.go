package dtos

import "github.com/shopspring/decimal"

type AccumulatePointsDTO struct {
	Id       int             `json:"user:id", validate:"required"`
	Sucursal int             `json:"sucursal_id", validate:"required"`
	Total    decimal.Decimal `json:"total_paid", validate:"required"`
}

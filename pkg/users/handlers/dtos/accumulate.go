package dtos

import (
	"github.com/leal/pkg/users/domain"
)

type AccumulatePointsDTO struct {
	Id       int                `json:"user_id", validate:"required"`
	Name     string             `json:"name", validate:"required"`
	Sucursal int                `json:"sucursal_id", validate:"required"`
	Payment  domain.UserPayment `json:"payment", validate:"required"`
}

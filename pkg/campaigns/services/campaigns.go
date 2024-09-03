package services

import (
	"context"

	c "github.com/leal/pkg/campaigns/domain"
	u "github.com/leal/pkg/users/domain"
	"github.com/shopspring/decimal"
)

// all different campaigns logic goes here
func CampaignTexaco1(ctx context.Context, name string, status c.CampaignStatus, payment u.UserPayment) float32 {
	if TEXACO_1 == name {
		if payment.PayDate.After(status.ActiveDate) && payment.PayDate.Before(status.InactiveDate) {
			return 2
		}
	}

	return 0
}

func CampaignTexaco2(ctx context.Context, name string, status c.CampaignStatus, payment u.UserPayment) float32 {
	if TEXACO_2 == name {
		if payment.Amount.GreaterThan(decimal.NewFromFloat32(20.000)) &&
			payment.PayDate.After(status.ActiveDate) && payment.PayDate.Before(status.InactiveDate) {
			return 0.3
		}
	}

	return 0
}

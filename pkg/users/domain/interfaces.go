package domain

import (
	"context"
)

type UserRepository interface {
	AccumulatePoints(ctx context.Context, name string, userId int, promotion float32) (*UserAccumulated, error)
	GetCampaignName(ctx context.Context, branchId int) (string, error)
	UpdateUser(ctx context.Context, id int, name string, acc *UserAccumulated) (*User, error)
}

type UserService interface {
	AccumulatePoints(ctx context.Context, name string, userId int, payment UserPayment) (*UserAccumulated, error)
}

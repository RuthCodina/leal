package domain

import (
	"context"
)

type CampaignRepository interface {
	Create(ctx context.Context, name string, description string, branchOffice int) (*Campaign, error)
	Activate(ctx context.Context, name string, id int) error
	Inactivate(ctx context.Context, name string, id int) error
	GetAllByBranch(ctx context.Context, branchId int, active bool) (*[]Campaign, error)
	GetAllByCommerce(ctx context.Context, commerceId int, active bool) (*[]Campaign, error)
}

type CampaignService interface {
	Create(ctx context.Context, name string, description string, branchOffice int) (*Campaign, error)
	Activate(ctx context.Context, name string, id int) error
	Inactivate(ctx context.Context, name string, id int) error
	GetAllByBranch(ctx context.Context, branchId int, active bool) (*[]Campaign, error)
	GetAllByCommerce(ctx context.Context, commerceId int, active bool) (*[]Campaign, error)
}

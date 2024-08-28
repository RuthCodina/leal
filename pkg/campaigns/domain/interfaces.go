package domain

import "context"

type CampaignRepository interface {
	Create(ctx context.Context, name string) (*Campaign, error)
	Activate(ctx context.Context, name string, id int) error
	Inactivate(ctx context.Context, name string, id int) error
	GetAllByBranch(ctx context.Context, name string, branchId int, active bool) (*[]Campaign, error)
	GetAllByCommerce(ctx context.Context, name string, commerceId int, active bool) (*[]Campaign, error)
}

type CampaignService interface {
	Create(ctx context.Context, name string) (*Campaign, error)
	Activate(ctx context.Context, name string, id int) error
	Inactivate(ctx context.Context, name string, id int) error
	GetAllByBranch(ctx context.Context, name string, branchId int, active bool) (*[]Campaign, error)
	GetAllByCommerce(ctx context.Context, name string, commerceId int, active bool) (*[]Campaign, error)
}

type CampaignHandler interface {
	Create(ctx context.Context, name string) (*Campaign, error)
	Activate(ctx context.Context, name string, id int) error
	Inactivate(ctx context.Context, name string, id int) error
	GetAllByBranch(ctx context.Context, name string, branchId int, active bool) (*[]Campaign, error)
	GetAllByCommerce(ctx context.Context, name string, commerceId int, active bool) (*[]Campaign, error)
}

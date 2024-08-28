package repository

import (
	"context"

	"github.com/leal/pkg/campaigns/domain"
)

func (r *campaignRepository) GetAllByBranch(ctx context.Context, name string, branchId int, active bool) (*[]domain.Campaign, error) {
	return nil, nil
}

func (r *campaignRepository) GetAllByCommerce(ctx context.Context, name string, commerceId int, active bool) (*[]domain.Campaign, error) {
	return nil, nil
}

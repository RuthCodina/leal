package repository

import (
	"context"

	"github.com/leal/pkg/campaigns/domain"
)

func (r *CampaignRepository) GetAllByBranch(ctx context.Context, name string, branchId int, active bool) (*[]domain.Campaign, error) {
	return nil, nil
}

func (r *CampaignRepository) GetAllByCommerce(ctx context.Context, name string, commerceId int, active bool) (*[]domain.Campaign, error) {
	return nil, nil
}

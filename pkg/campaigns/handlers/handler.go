package handlers

import (
	"context"
	"net/http"

	"github.com/leal/pkg/campaigns/domain"
)

type Handler struct {
	CampaignService domain.CampaignService
}

type CampaignHandlers interface {
	Create(ctx context.Context, w http.ResponseWriter, r *http.Request)
	Activate(ctx context.Context, name string, id int) error
	Inactivate(ctx context.Context, name string, id int) error
	GetAllByBranch(ctx context.Context, name string, branchId int, active bool)
	GetAllByCommerce(ctx context.Context, name string, commerceId int, active bool)
}

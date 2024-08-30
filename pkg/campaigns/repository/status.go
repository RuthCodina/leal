package repository

import (
	"context"
	"database/sql"

	"github.com/leal/pkg/campaigns/domain"
	"github.com/leal/pkg/helpers"
)

func (r *Repository) Status(ctx context.Context, branchId int) (*domain.CampaignStatus, error) {
	campaignStatusDB := &domain.CampaignStatusDB{}
	campaignStatus := &domain.CampaignStatus{}
	if err := r.db.QueryRow(campaignStatusQuery, branchId).Scan(&campaignStatusDB.Id, &campaignStatusDB.Name, &campaignStatusDB.ActiveDate, &campaignStatusDB.InactiveDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.ErrAccumulation
		}
		return nil, err
	}
	campaignStatus = &domain.CampaignStatus{
		Id:           campaignStatusDB.Id,
		Name:         campaignStatusDB.Name,
		ActiveDate:   campaignStatusDB.ActiveDate.Time,
		InactiveDate: campaignStatusDB.ActiveDate.Time,
	}
	return campaignStatus, nil
}

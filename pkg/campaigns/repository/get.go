package repository

import (
	"context"
	"log"

	"github.com/leal/pkg/campaigns/domain"
)

func (r *Repository) GetAllByBranch(ctx context.Context, branchId int, active bool) (*[]domain.Campaign, error) {
	campaigns := []domain.Campaign{}
	act := 0

	if active {
		act = 1
	}

	rows, err := r.db.Query(getByBranchQuery, branchId, act)
	if err != nil {
		log.Printf("cannot excute select query: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		campaignDB := domain.CampaignSQL{}
		err := rows.Scan(&campaignDB.Id, &campaignDB.Name, &campaignDB.ActiveDate, &campaignDB.InactiveDate, &campaignDB.State, &campaignDB.BranchOffice, &campaignDB.Description, &campaignDB.CreatedAt, &campaignDB.UpdatedAt)
		if err != nil {
			log.Println("cannot read current row")
			return nil, err
		}
		campaign := domain.Campaign{
			Id:           campaignDB.Id,
			Name:         campaignDB.Name,
			ActiveDate:   campaignDB.ActiveDate.Time,
			InactiveDate: campaignDB.ActiveDate.Time,
			State:        campaignDB.State,
			BranchOffice: campaignDB.BranchOffice,
			Description:  campaignDB.Description,
			CreatedAt:    campaignDB.CreatedAt.Time,
			UpdatedAt:    campaignDB.UpdatedAt.Time,
		}
		campaigns = append(campaigns, campaign)
	}

	return &campaigns, nil
}

func (r *Repository) GetAllByCommerce(ctx context.Context, commerceId int, active bool) (*[]domain.Campaign, error) {
	campaigns := []domain.Campaign{}
	act := 0

	if active {
		act = 1
	}
	rows, err := r.db.Query(getByCommerceQuery, commerceId, act)
	if err != nil {
		log.Printf("cannot excute select query: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		campaignDB := domain.CampaignSQL{}
		err := rows.Scan(&campaignDB.Id, &campaignDB.Name, &campaignDB.ActiveDate, &campaignDB.InactiveDate, &campaignDB.State, &campaignDB.BranchOffice, &campaignDB.Description, &campaignDB.CreatedAt, &campaignDB.UpdatedAt)
		if err != nil {
			log.Println("cannot read current row")
			return nil, err
		}
		campaign := domain.Campaign{
			Id:           campaignDB.Id,
			Name:         campaignDB.Name,
			ActiveDate:   campaignDB.ActiveDate.Time,
			InactiveDate: campaignDB.ActiveDate.Time,
			State:        campaignDB.State,
			BranchOffice: campaignDB.BranchOffice,
			Description:  campaignDB.Description,
			CreatedAt:    campaignDB.CreatedAt.Time,
			UpdatedAt:    campaignDB.UpdatedAt.Time,
		}
		campaigns = append(campaigns, campaign)
	}

	return &campaigns, nil
}

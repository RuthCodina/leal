package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/leal/pkg/helpers"
)

func (r *Repository) Activate(ctx context.Context, name string, id int) error {
	var branchActivated bool
	if err := r.db.QueryRow(campaignActiveInBranch, id).Scan(&branchActivated); err != nil {
		if err == sql.ErrNoRows {
			_, err := r.db.Exec(activateCampaignQuery, name, id)
			if err != nil {
				log.Printf("cannot activate the campaign: %s", err.Error())
				return err
			}
			_, err = r.db.Exec(branchCampaignActivate, id)
			if err != nil {
				log.Printf("cannot activate the campaign in sucursal table: %s", err.Error())
				return err
			}
			_, err = r.db.Exec(updateCommerceIdInBranch, id, id)
			if err != nil {
				log.Printf("cannot update the campaign_id in sucursal table: %s", err.Error())
				return err
			}
			return nil
		}
		return err
	} else if branchActivated {
		return helpers.ErrActivateCampaign
	}
	return nil
}
func (r *Repository) Inactivate(ctx context.Context, name string, id int) error {
	_, err := r.db.Exec(inactivateCampaignQuery, name, id)
	if err != nil {
		log.Printf("cannot inactivate the campaign: %s", err.Error())
		return err
	}
	_, err = r.db.Exec(branchCampaignInactivate, id)
	if err != nil {
		log.Printf("cannot inactivate the campaign: %s", err.Error())
		return err
	}
	return nil
}

func (r *Repository) SetActivate(ctx context.Context, name string, id int) error {
	return nil
}
func (r *Repository) SetInactivate(ctx context.Context, name string, id int) error {
	return nil
}

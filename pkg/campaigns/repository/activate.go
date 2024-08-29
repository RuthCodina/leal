package repository

import (
	"context"
	"log"
)

func (r *Repository) Activate(ctx context.Context, name string, id int) error {
	_, err := r.db.Exec(activateCampaignQuery, name, id)
	if err != nil {
		log.Printf("cannot activate the campaign: %s", err.Error())
		return err
	}
	return nil
}
func (r *Repository) Inactivate(ctx context.Context, name string, id int) error {
	_, err := r.db.Exec(inactivateCampaignQuery, name, id)
	if err != nil {
		log.Printf("cannot inactivate the campaign: %s", err.Error())
		return err
	}
	return nil
}

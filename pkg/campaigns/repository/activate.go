package repository

import "context"

func (r *CampaignRepository) Activate(ctx context.Context, name string, id int) error {
	return nil
}
func (r *CampaignRepository) Inactivate(ctx context.Context, name string, id int) error {
	return nil
}

package repository

import (
	"database/sql"

	"github.com/leal/pkg/campaigns/domain"
)

var _ domain.CampaignRepository = &CampaignRepository{}

type CampaignRepository struct {
	Db *sql.DB
}

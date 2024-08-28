package repository

import (
	"database/sql"

	"github.com/leal/pkg/campaigns/domain"
)

var _ domain.CampaignRepository = &campaignRepository{}

type campaignRepository struct {
	db *sql.DB
}

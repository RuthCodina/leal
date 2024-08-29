package repository

import (
	"database/sql"

	"github.com/leal/pkg/campaigns/domain"
)

var _ domain.CampaignRepository = &Repository{}

type Repository struct {
	db *sql.DB
}

func NewCampaingRepository(db *sql.DB) *Repository {
	return &Repository{
		db,
	}
}

package domain

import "time"

type Campaign struct {
	Id           int
	Name         string
	ActiveDate   time.Time
	InactiveDate time.Time
	State        bool
	BranchOffice int
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Branch struct {
	Id            int
	Name          string
	CommerceId    int
	CampaignId    int
	CampaignState bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

package domain

import "time"

type Campaign struct {
	id           int
	name         string
	activeDate   time.Time
	inactiveDate time.Time
	state        bool
	branchOffice int
	description  string
}

type Branch struct {
	id            int
	name          string
	commerceId    int
	campaignId    int
	campaignState bool
}

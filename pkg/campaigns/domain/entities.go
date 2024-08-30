package domain

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

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

type CampaignSQL struct {
	Id           int
	Name         string
	ActiveDate   NullTime
	InactiveDate NullTime
	State        bool
	BranchOffice int
	Description  string
	CreatedAt    NullTime
	UpdatedAt    NullTime
}

type CampaignStatusDB struct {
	Id           int
	Name         string
	ActiveDate   NullTime
	InactiveDate NullTime
}
type CampaignStatus struct {
	Id           int
	Name         string
	ActiveDate   time.Time
	InactiveDate time.Time
}

type NullTime struct {
	mysql.NullTime
}

func (nt *NullTime) MarshallJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}

package domain

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
)

type User struct {
	Id        int
	Name      string
	LastName  string
	DNI       int
	State     bool
	Points    decimal.Decimal
	LealCoins decimal.Decimal
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserPayment struct {
	Amount   decimal.Decimal
	PayDate  time.Time
	BranchId int
}

type UserAccumulated struct {
	Cashback decimal.Decimal
	Points   decimal.Decimal
}

type CampaignBranch struct {
	BranchId     int
	CampaignName string
}

type UserDB struct {
	Id        int
	Name      string
	LastName  string
	DNI       int
	State     bool
	Points    decimal.Decimal
	LealCoins decimal.Decimal
	CreatedAt NullTime
	UpdatedAt NullTime
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

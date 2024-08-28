package domain

import "github.com/shopspring/decimal"

type User struct {
	id        int
	name      string
	lastName  string
	dni       int
	state     bool
	points    decimal.Decimal
	lealCoins decimal.Decimal
}

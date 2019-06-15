package entity

import "time"

type Balance struct {
	BalanceId int
	Amount int
	Item string
	KindId int
	PurposeId int
	PlaceId int
	Date time.Time
}

type Balances []Balance

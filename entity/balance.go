package entity

import "time"

type Balance struct {
	BalanceId int `json:"balance_id"`
	Amount int `json:"amount"`
	Item string `json:"item"`
	KindId int `json:"kind_id`
	PurposeId int `json:"purpose_id"`
	PlaceId int `json:place_id"`
	Date time.Time `json:"date"`
}

type Balances []Balance

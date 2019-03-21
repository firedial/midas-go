package model

import (
    "github.com/firedial/go-midas/db"
)

type Balance struct {
    Id int `json:"id"`
    Amount int `json:"amount"`
    Item string `json:"item"`
    Kind_id int `json:"kind_id"`
    Purpose_id int `json:"purpose_id"`
    Place_id int `json:"place_id"`
    Date string `json:"date"`
}

func GetBalance(id string) Balance {
    var balance Balance

    db := db.Init();
    defer db.Close();

    where := "";
    where = "WHERE balance_id = " + string(id)

    err := db.QueryRow(`
        SELECT
            balance_id,
            amount,
            item,
            kind_id,
            purpose_id,
            place_id,
            date
        FROM
            balance
    ` + where).Scan(
        &(balance.Id),
        &(balance.Amount),
        &(balance.Item),
        &(balance.Kind_id),
        &(balance.Purpose_id),
        &(balance.Place_id),
        &(balance.Date))

  if err != nil {
    panic(err.Error())
  }
  return balance
}

func InsertBalance(balance Balance) bool {
    db := db.Init();
    defer db.Close();

    stmt, err := db.Prepare(`
        INSERT INTO balance (
            amount,
            item,
            kind_id,
            purpose_id,
            place_id,
            date
        ) VALUES 
        (?, ?, ?, ?, ?, ?)`)

  if err != nil {
    panic(err.Error())
  }
  defer stmt.Close()

  _, err = stmt.Exec(balance.Amount, balance.Item, balance.Kind_id, balance.Purpose_id, balance.Place_id, balance.Date)
  if err != nil {
    panic(err.Error())
  }

  return true
}

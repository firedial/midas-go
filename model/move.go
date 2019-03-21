package model

import (
    "github.com/firedial/midas-go/db"
)

type Move struct {
    Attribute string `json:"attribute"`
    Amount int `json:"amount"`
    BeforeId int `json:"before_id"`
    AfterId int `json:"after_id"`
    Date string `json:"date"`
}

func InsertMove(move Move) bool {
    db := db.Init();
    defer db.Close();

    tx, err := db.Begin()
    if err != nil {
        panic(err)
    }

    defer func() {
        err := recover()
        if err != nil {
            tx.Rollback()
        }
    }()

    beforeBalance, afterBalance := getMoveBalance(move)

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

    _, err = stmt.Exec(beforeBalance.Amount, beforeBalance.Item, beforeBalance.Kind_id, beforeBalance.Purpose_id, beforeBalance.Place_id, beforeBalance.Date)
    _, err = stmt.Exec(afterBalance.Amount, afterBalance.Item, afterBalance.Kind_id, afterBalance.Purpose_id, afterBalance.Place_id, afterBalance.Date)
    if err != nil {
        panic(err.Error())
    }

    tx.Commit()

    return true
}

func getMoveBalance(move Move) (Balance, Balance) {
    var beforeBalance Balance
    var afterBalance Balance

    beforeBalance.Amount = -1 * move.Amount
    afterBalance.Amount = move.Amount
    beforeBalance.Kind_id = 14
    afterBalance.Kind_id = 14
    beforeBalance.Date = move.Date
    afterBalance.Date = move.Date

    if move.Attribute == "purpose" {
        beforeBalance.Item = "予算移動元"
        afterBalance.Item = "予算移動先"
        beforeBalance.Place_id = 4
        afterBalance.Place_id = 4

        beforeBalance.Purpose_id = move.BeforeId 
        afterBalance.Purpose_id = move.AfterId
    } else if move.Attribute == "place" {
        beforeBalance.Item = "場所移動元"
        afterBalance.Item = "場所移動先"
        beforeBalance.Purpose_id = 12 
        afterBalance.Purpose_id = 12 

        beforeBalance.Place_id = move.BeforeId 
        afterBalance.Place_id = move.AfterId
    } else {
        panic("non")
    }

    return beforeBalance, afterBalance
}

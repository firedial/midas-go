package dao

import (
    "time"

    "github.com/firedial/midas-go/entity"
    "github.com/firedial/midas-go/db"
)

type MysqlBalanceRepository struct {
    
}

func (MysqlBalanceRepository) FindAll() (balances entity.Balances, err error) {
    db := db.Init();
    defer db.Close();

    //rows, err := db.Query("SELECT balance_id, amount, item, kind_id, purpose_id, place_id, date FROM balances")
    rows, err := db.Query("SELECT * FROM balance")
    defer rows.Close()
    if err != nil {
        return
    }
    for rows.Next() {
        var balance_id int
        var amount int
        var item string
        var kind_id int
        var purpose_id int
        var place_id int
        var date time.Time

        err := rows.Scan(&balance_id, &amount, &item, &kind_id, &purpose_id, &place_id, &date)
        if err != nil {
            continue
        }
        balance := entity.Balance{
            BalanceId: balance_id,
            Amount: amount,
            Item: item,
            KindId: kind_id,
            PurposeId: purpose_id,
            PlaceId: place_id,
            Date: date,
        }
        balances = append(balances, balance)
    }
    return
}

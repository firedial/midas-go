package controller

import(
    "github.com/firedial/midas-go/model"
    "github.com/firedial/midas-go/entity"
)

func BalanceGet(queries map[string][]string) entity.Balances {
    id := queries["id"][0]
    balance := model.GetBalance(id)
    return balance
}

func BalancePost(balance model.Balance) model.Balance {
    _ = model.InsertBalance(balance)
    return balance 
}

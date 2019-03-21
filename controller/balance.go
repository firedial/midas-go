package controller

import(
    "github.com/firedial/midas-go/model"
)

func BalanceGet(queries map[string][]string) model.Balance {
    id := queries["id"][0]
    balance := model.GetBalance(id)
    return balance
}

func BalancePost(balance model.Balance) model.Balance {
    _ = model.InsertBalance(balance)
    return balance 
}

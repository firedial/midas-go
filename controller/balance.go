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
    /*
    balance.Amount, _ = strconv.Atoi(post_forms["amount"])
    balance.Item = post_forms["item"]
    balance.Kind_id, _ = strconv.Atoi(post_forms["kind_id"])
    balance.Purpose_id, _ = strconv.Atoi(post_forms["purpose_id"])
    balance.Place_id, _ = strconv.Atoi(post_forms["place_id"])
    balance.Date = post_forms["date"]
    */
    _ = model.InsertBalance(balance)
    return balance 
}

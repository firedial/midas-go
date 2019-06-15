package controller

import(
    "github.com/firedial/midas-go/entity"
    "github.com/firedial/midas-go/interactor"
)

func BalanceGet(queries map[string][]string) entity.Balances {
    //id := queries["id"][0]
    balances, _ := interactor.GetBalance()
    return balances
}

func BalancePost(balances entity.Balances) string {
    rc := interactor.InsertBalances(balances)
    return rc
}

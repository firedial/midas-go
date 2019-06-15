package interactor

import (
    "github.com/firedial/midas-go/entity"
    "github.com/firedial/midas-go/repository"
    "github.com/firedial/midas-go/dao"
)

var balanceRepository repository.BalanceRepository = &dao.MysqlBalanceRepository{}

func GetBalance() (entity.Balances, error) {
    return balanceRepository.FindAll()
}

func InsertBalances(balances entity.Balances) string {
    _ = balanceRepository.SaveAll(balances);
    return "aaa"
}

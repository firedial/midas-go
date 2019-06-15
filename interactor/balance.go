package interactor

import (
    "github.com/firedial/midas-go/entity"
    "github.com/firedial/midas-go/repository"
    "github.com/firedial/midas-go/dao"
)


func GetBalancea() entity.Balances {
    var balanceDao repository.BalanceRepository = &dao.MysqlBalanceRepository{}
    balances, _ := balanceDao.FindAll();
    return balances
}

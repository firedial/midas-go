package repository

import "github.com/firedial/midas-go/entity"

type BalanceRepository interface {
    //Save(entity.Balance) error
    //SaveAll(entity.Balances) error
    //FindById(int) (entity.Balance, error)
    FindAll() (entity.Balances, error)
}

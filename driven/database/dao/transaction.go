package dao

import dto "github.com/devbenatti/rbwallet-go/dto/transaction"

var _ TransactionDAO = (*Transaction)(nil)

type Transaction struct {
}

func (td *Transaction) Create(t dto.TransactionDTO) {

}

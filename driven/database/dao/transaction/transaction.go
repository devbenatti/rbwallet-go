package dao

import (
	"database/sql"

	dto "github.com/devbenatti/rbwallet-go/dto/transaction"
)

var _ TransactionDAO = (*Transaction)(nil)

type Transaction struct {
	db *sql.DB
}

func NewTransactionDAO(conn *sql.DB) Transaction {
	return Transaction{
		db: conn,
	}
}

func (td *Transaction) Create(t dto.TransactionDTO) {

}

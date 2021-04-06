package dao

import (
	"database/sql"

	"github.com/devbenatti/rbwallet-go/model"
)

var _ TransactionDAO = (*Transaction)(nil)

type Transaction struct {
	db *sql.DB
}

func NewTransactionDAO(conn *sql.DB) *Transaction {
	return &Transaction{
		db: conn,
	}
}

func (td *Transaction) Create(t model.Transaction) {
	defer td.db.Close()

	insert, err := td.db.Prepare("INSERT INTO transactions(code, account_id, operation_id, total) VALUES(?,?,?,?)")

	if err != nil {
		panic(err)
	}

	_, err = insert.Exec(t.Code().String(), t.AccountID().String(), t.Operation().Type(), t.Total().String())

	if err != nil {
		panic(err)
	}
}

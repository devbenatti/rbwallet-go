package model

import (
	"github.com/devbenatti/rbwallet-go/model/valueObject"
)

type Transaction struct {
	code      valueObject.Uuid
	accountID valueObject.Uuid
	operation Operation
	total     *valueObject.Money
}

func NewTransaction(c string, accountID string, operation int, total float64) Transaction {
	return Transaction{
		code:      valueObject.NewUuid(c),
		accountID: valueObject.NewUuid(accountID),
		operation: NewOperation(operation),
		total:     valueObject.NewMoney(total),
	}
}

func (t *Transaction) Code() *valueObject.Uuid {
	return &t.code
}

func (t *Transaction) AccountID() *valueObject.Uuid {
	return &t.accountID
}

func (t *Transaction) Operation() *Operation {
	return &t.operation
}

func (t *Transaction) Total() *valueObject.Money {
	return t.total
}

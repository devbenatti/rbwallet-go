package dto

import (
	"github.com/devbenatti/rbwallet-go/model"
	"github.com/devbenatti/rbwallet-go/model/valueObject"
)

type Transaction struct {
	code      valueObject.Uuid
	accountID int
	operation model.Operation
	total     valueObject.Money
}

func NewTransaction(c valueObject.Uuid, a int, o int, t float64) Transaction {
	return Transaction{
		code:      c,
		accountID: a,
		operation: model.NewOperation(o),
		total:     valueObject.NewMoney(t),
	}
}

func (t *Transaction) Code() valueObject.Uuid {
	return t.code
}

func (t *Transaction) AccountID() int {
	return t.accountID
}

func (t *Transaction) Operation() *model.Operation {
	return &t.operation
}

func (t *Transaction) Total() *valueObject.Money {
	return &t.total
}

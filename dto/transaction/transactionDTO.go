package dto

import (
	"github.com/devbenatti/rbwallet-go/model"
	"github.com/devbenatti/rbwallet-go/model/valueObject"
)

type TransactionDTO struct {
	code      valueObject.Uuid
	accountID int
	operation model.Operation
	total     *valueObject.Money
}

func NewTransactionDTO(c valueObject.Uuid, accountID int, operation int, total float64) TransactionDTO {
	return TransactionDTO{
		code:      c,
		accountID: accountID,
		operation: model.NewOperation(operation),
		total:     valueObject.NewMoney(total),
	}
}

func (t *TransactionDTO) Code() valueObject.Uuid {
	return t.code
}

func (t *TransactionDTO) AccountID() int {
	return t.accountID
}

func (t *TransactionDTO) Operation() *model.Operation {
	return &t.operation
}

func (t *TransactionDTO) Total() *valueObject.Money {
	return t.total
}

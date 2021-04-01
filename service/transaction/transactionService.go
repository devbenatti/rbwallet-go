package transaction

import (
	"github.com/devbenatti/rbwallet-go/dto"
	"github.com/devbenatti/rbwallet-go/model/valueObject"
)

type TransactionService struct {
}

func (ts *TransactionService) Create(t dto.Transaction) valueObject.Uuid {

	if t.Operation().Flow().IsOutFlow() {
		t.Total().Multiply(-1)
	}

	return t.Code()
}

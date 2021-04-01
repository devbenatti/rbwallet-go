package transaction

import (
	"github.com/devbenatti/rbwallet-go/dto"
)

type TransactionService struct {
}

func (ts *TransactionService) Create(t dto.Transaction) {

	if t.Operation().Flow().IsOutFlow() {
		t.Total().Multiply(-1)
	}
}

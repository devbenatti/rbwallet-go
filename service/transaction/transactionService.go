package service

import (
	"github.com/devbenatti/rbwallet-go/driven/database/dao"
	dto "github.com/devbenatti/rbwallet-go/dto/transaction"
)

type TransactionService struct {
	transactionDAO dao.TransactionDAO
}

func NewTransactionService(td dao.TransactionDAO) TransactionService {
	return TransactionService{
		transactionDAO: td,
	}
}

func (ts *TransactionService) Create(t dto.TransactionDTO) {

	if t.Operation().Flow().IsOutFlow() {
		t.Total().Negative()
	}

	ts.transactionDAO.Create(t)
}

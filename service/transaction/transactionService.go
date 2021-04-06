package service

import (
	dao "github.com/devbenatti/rbwallet-go/driven/database/dao/transaction"
	dto "github.com/devbenatti/rbwallet-go/dto/transaction"
	"github.com/devbenatti/rbwallet-go/model"
)

type TransactionService struct {
	transactionDAO dao.TransactionDAO
}

func NewTransactionService(td dao.TransactionDAO) TransactionService {
	return TransactionService{
		transactionDAO: td,
	}
}

func (ts *TransactionService) Create(ct dto.CreateTransactionDTO) {

	t := model.NewTransaction(ct.Code, ct.AccountID, ct.Operation, ct.Total)

	if t.Operation().Flow().IsOutFlow() {
		t.Total().Negative()
	}

	ts.transactionDAO.Create(t)
}

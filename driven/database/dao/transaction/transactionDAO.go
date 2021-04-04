package dao

import dto "github.com/devbenatti/rbwallet-go/dto/transaction"

type TransactionDAO interface {
	Create(t dto.TransactionDTO)
}

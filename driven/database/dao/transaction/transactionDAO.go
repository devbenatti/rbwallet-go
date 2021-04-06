package dao

import "github.com/devbenatti/rbwallet-go/model"

type TransactionDAO interface {
	Create(t model.Transaction)
}

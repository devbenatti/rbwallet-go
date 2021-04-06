package action

import (
	"github.com/devbenatti/rbwallet-go/driven/database"
	dao "github.com/devbenatti/rbwallet-go/driven/database/dao/transaction"
	"github.com/devbenatti/rbwallet-go/driven/uuid"
	service "github.com/devbenatti/rbwallet-go/service/transaction"
)

type transactionBehavior struct {
	uuid    uuid.UuidGenerator
	dao     dao.TransactionDAO
	service service.TransactionService
}

func (ac *transactionBehavior) configure() {
	ac.uuid = uuid.NewUuidGenerator()
	db := database.GetConnection()
	ac.dao = dao.NewTransactionDAO(db)
	ac.service = service.NewTransactionService(ac.dao)
}

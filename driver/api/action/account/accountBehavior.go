package action

import (
	"github.com/devbenatti/rbwallet-go/driven/database"
	dao "github.com/devbenatti/rbwallet-go/driven/database/dao/account"
	"github.com/devbenatti/rbwallet-go/driven/uuid"
	service "github.com/devbenatti/rbwallet-go/service/account"
)

type accountBehavior struct {
	uuid    uuid.UuidGenerator
	dao     dao.AccountDAO
	service service.AccountService
}

func (ac *accountBehavior) configure() {
	ac.uuid = uuid.NewUuidGenerator()
	db := database.GetConnection()
	ac.dao = dao.NewAccountDAO(db)
	ac.service = service.NewAccountService(ac.dao)
}

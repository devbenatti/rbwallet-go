package service

import (
	dao "github.com/devbenatti/rbwallet-go/driven/database/dao/account"
	dto "github.com/devbenatti/rbwallet-go/dto/account"
	"github.com/devbenatti/rbwallet-go/model/valueObject"
)

type AccountService struct {
	accountDAO dao.AccountDAO
}

func NewAccountService(ad dao.AccountDAO) AccountService {
	return AccountService{
		accountDAO: ad,
	}
}

func (as *AccountService) Create(a dto.AccountDTO) {
	as.accountDAO.Create(a)
}

func (as *AccountService) FindOne(id valueObject.Uuid) *dto.AccountDTO {
	return as.accountDAO.FindByID(id)
}

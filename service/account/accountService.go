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
	err := as.accountDAO.Create(a)

	if err != nil {

	}

}

func (as *AccountService) FindOne(id valueObject.Uuid) (*dto.AccountDTO, error) {
	account, err := as.accountDAO.FindByID(id)

	if err != nil {
		return nil, err
	}

	return account, nil
}

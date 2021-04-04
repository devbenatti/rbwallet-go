package service

import (
	dao "github.com/devbenatti/rbwallet-go/driven/database/dao/account"
	dto "github.com/devbenatti/rbwallet-go/dto/account"
	model "github.com/devbenatti/rbwallet-go/model/errors"
	"github.com/devbenatti/rbwallet-go/model/valueObject"
)

type AccountNotFoundErr = model.ErrEntityNotFound

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

func (as *AccountService) FindOne(id valueObject.Uuid) (*dto.AccountDTO, error) {

	acc := as.accountDAO.FindByID(id)

	if acc == nil {
		return nil, model.NewErrEntityNotFound("account not found")
	}

	return acc, nil
}

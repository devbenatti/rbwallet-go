package service

import (
	"reflect"
	"testing"

	dao "github.com/devbenatti/rbwallet-go/driven/database/dao/account"
	dto "github.com/devbenatti/rbwallet-go/dto/account"
	"github.com/devbenatti/rbwallet-go/model/valueObject"
)

var _ dao.AccountDAO = (*AccountDAOSpy)(nil)

type AccountDAOSpy struct {
	receivedDTO dto.AccountDTO
}

func (ad *AccountDAOSpy) Create(a dto.AccountDTO) {
	ad.receivedDTO = a
}

func (ad *AccountDAOSpy) FindByID(id valueObject.Uuid) *dto.AccountDTO {
	acc := dto.NewAccountDTO("59747936-f126-4286-a214-b15b9d9754e5", "12345678912")
	return &acc
}

func TestAccountServiceCreate(t *testing.T) {
	dao := &AccountDAOSpy{}

	ac := dto.NewAccountDTO("59747936-f126-4286-a214-b15b9d9754e5", "12345678912")
	expectedDTO := dto.NewAccountDTO("59747936-f126-4286-a214-b15b9d9754e5", "12345678912")

	s := NewAccountService(dao)

	s.Create(ac)

	if !reflect.DeepEqual(dao.receivedDTO, expectedDTO) {
		t.Errorf("expect '%v' - result '%v'", expectedDTO, dao.receivedDTO)
	}
}

func TestAccountServiceFindByID(t *testing.T) {
	dao := &AccountDAOSpy{}

	id := valueObject.NewUuid("59747936-f126-4286-a214-b15b9d9754e5")

	expectedDTO := dto.NewAccountDTO("59747936-f126-4286-a214-b15b9d9754e5", "12345678912")

	s := NewAccountService(dao)

	acc, _ := s.FindOne(id)

	if !reflect.DeepEqual(acc, &expectedDTO) {
		t.Errorf("expect '%v' - result '%v'", expectedDTO, acc)
	}
}

var _ dao.AccountDAO = (*AccountDAOSpyTwo)(nil)

type AccountDAOSpyTwo struct {
	receivedDTO dto.AccountDTO
}

func (ad *AccountDAOSpyTwo) Create(a dto.AccountDTO) {
	ad.receivedDTO = a
}

func (ad *AccountDAOSpyTwo) FindByID(id valueObject.Uuid) *dto.AccountDTO {
	return nil
}

func TestAccountServiceFindByIDShouldBeReturnErr(t *testing.T) {
	dao := &AccountDAOSpyTwo{}

	id := valueObject.NewUuid("59747936-f126-4286-a214-b15b9d9754e5")

	s := NewAccountService(dao)

	_, err := s.FindOne(id)

	if _, ok := err.(*AccountNotFoundErr); !ok {
		t.Errorf("error should be type of AccountNotFoundErr")
	}
}

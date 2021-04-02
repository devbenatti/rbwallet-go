package service

import (
	"reflect"
	"testing"

	"github.com/devbenatti/rbwallet-go/driven/database/dao"
	dto "github.com/devbenatti/rbwallet-go/dto/transaction"
	"github.com/devbenatti/rbwallet-go/model/valueObject"
)

var _ dao.TransactionDAO = (*TransactionDAOSpy)(nil)

type TransactionDAOSpy struct {
	receivedDTO dto.TransactionDTO
}

func (ts *TransactionDAOSpy) Create(t dto.TransactionDTO) {
	ts.receivedDTO = t
}

func TestCreate(t *testing.T) {
	dataProvider := []struct {
		dto         dto.TransactionDTO
		expectedDTO dto.TransactionDTO
	}{
		{
			dto: dto.NewTransactionDTO(
				valueObject.NewUuid("59747936-f126-4286-a214-b15b9d9754e5"),
				1,
				1,
				10.00,
			),
			expectedDTO: dto.NewTransactionDTO(
				valueObject.NewUuid("59747936-f126-4286-a214-b15b9d9754e5"),
				1,
				1,
				-10.00,
			),
		},
		{
			dto: dto.NewTransactionDTO(
				valueObject.NewUuid("59747936-f126-4286-a214-b15b9d9754e5"),
				1,
				2,
				11.00,
			),
			expectedDTO: dto.NewTransactionDTO(
				valueObject.NewUuid("59747936-f126-4286-a214-b15b9d9754e5"),
				1,
				2,
				-11.00,
			),
		},
		{
			dto: dto.NewTransactionDTO(
				valueObject.NewUuid("59747936-f126-4286-a214-b15b9d9754e5"),
				1,
				4,
				12.00,
			),
			expectedDTO: dto.NewTransactionDTO(
				valueObject.NewUuid("59747936-f126-4286-a214-b15b9d9754e5"),
				1,
				4,
				12.00,
			),
		},
	}

	dao := &TransactionDAOSpy{}

	s := NewTransactionService(dao)

	for _, tt := range dataProvider {
		t.Run("Create transaction", func(t *testing.T) {
			s.Create(tt.dto)
			if !reflect.DeepEqual(dao.receivedDTO, tt.expectedDTO) {
				t.Errorf("expect '%v' - result '%v'", tt.expectedDTO, dao.receivedDTO)
			}
		})
	}

}

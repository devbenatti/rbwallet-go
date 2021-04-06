package service

import (
	"reflect"
	"testing"

	dao "github.com/devbenatti/rbwallet-go/driven/database/dao/transaction"
	dto "github.com/devbenatti/rbwallet-go/dto/transaction"
	"github.com/devbenatti/rbwallet-go/model"
)

var _ dao.TransactionDAO = (*TransactionDAOSpy)(nil)

type TransactionDAOSpy struct {
	receivedTransaction model.Transaction
}

func (ts *TransactionDAOSpy) Create(t model.Transaction) {
	ts.receivedTransaction = t
}

func TestCreate(t *testing.T) {
	dataProvider := []struct {
		dto      dto.CreateTransactionDTO
		expected model.Transaction
	}{
		{
			dto: dto.CreateTransactionDTO{
				Code:      "59747936-f126-4286-a214-b15b9d9754e5",
				AccountID: "573F9ECC-703B-43D6-697B-8125FB389E46",
				Operation: 1,
				Total:     10.00,
			},
			expected: model.NewTransaction(
				"59747936-f126-4286-a214-b15b9d9754e5",
				"573F9ECC-703B-43D6-697B-8125FB389E46",
				1,
				-10.00,
			),
		},
		{
			dto: dto.CreateTransactionDTO{
				Code:      "59747936-f126-4286-a214-b15b9d9754e5",
				AccountID: "573F9ECC-703B-43D6-697B-8125FB389E46",
				Operation: 2,
				Total:     11.00,
			},
			expected: model.NewTransaction(
				"59747936-f126-4286-a214-b15b9d9754e5",
				"573F9ECC-703B-43D6-697B-8125FB389E46",
				2,
				-11.00,
			),
		},
		{
			dto: dto.CreateTransactionDTO{
				Code:      "59747936-f126-4286-a214-b15b9d9754e5",
				AccountID: "573F9ECC-703B-43D6-697B-8125FB389E46",
				Operation: 4,
				Total:     12.00,
			},
			expected: model.NewTransaction(
				"59747936-f126-4286-a214-b15b9d9754e5",
				"573F9ECC-703B-43D6-697B-8125FB389E46",
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
			if !reflect.DeepEqual(dao.receivedTransaction, tt.expected) {
				t.Errorf("expect '%v' - result '%v'", tt.expected, dao.receivedTransaction)
			}
		})
	}

}

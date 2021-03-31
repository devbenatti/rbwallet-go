package model

import "testing"

func TestNewTransaction(t *testing.T) {

	testsTransaction := []struct {
		operationType         int
		total                 float64
		expectedTotal         float64
		expectedOperationType int
	}{
		{
			operationType: 1,
			total:         50.00,
			expectedTotal: -50.00,
		},
		{
			operationType: 2,
			total:         10.00,
			expectedTotal: -10.00,
		},
		{
			operationType: 3,
			total:         11.00,
			expectedTotal: -11.00,
		},
		{
			operationType: 4,
			total:         10.00,
			expectedTotal: 10.00,
		},
	}

	for _, tt := range testsTransaction {
		t.Run("test create New Transaction", func(t *testing.T) {
			transaction := NewTransaction(tt.operationType, tt.total)
			operationType := transaction.Operation().Type()

			if int(operationType) != tt.operationType {
				t.Errorf("expect type '%d', result '%d'", tt.operationType, operationType)
			}

			if transaction.Total().Val() != tt.expectedTotal {
				t.Errorf("expect total '%.2f', result '%.2f'", tt.expectedTotal, transaction.Total().Val())
			}
		})
	}
}

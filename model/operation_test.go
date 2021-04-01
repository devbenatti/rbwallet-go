package model

import (
	"reflect"
	"testing"
)

func TestNewOperation(t *testing.T) {
	operationTests := []struct {
		operationID int
		expected    Operation
	}{
		{
			operationID: 1,
			expected: Operation{
				operationType: 1,
				flow:          1,
			},
		},
		{
			operationID: 2,
			expected: Operation{
				operationType: 2,
				flow:          1,
			},
		},
		{
			operationID: 3,
			expected: Operation{
				operationType: 3,
				flow:          1,
			},
		},
		{
			operationID: 4,
			expected: Operation{
				operationType: 4,
				flow:          2,
			},
		},
	}

	for _, tt := range operationTests {
		t.Run("New Operation", func(t *testing.T) {
			result := NewOperation(tt.operationID)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expect '%v' - result '%v'", tt.expected, result)
			}
		})
	}
}

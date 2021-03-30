package model

import "math"

const float64EqualityThreshold = 1e-9

type Amount struct {
	val float64
}

type Money struct {
	amount *Amount
}

func (m *Money) Amount() float64 {
	return m.amount.val
}

func (m *Money) Add(om *Money) *Money {
	return &Money{
		amount: &Amount{val: m.amount.val + om.amount.val},
	}
}

func (m *Money) Subtract(om *Money) *Money {
	return &Money{
		amount: &Amount{val: m.amount.val - om.amount.val},
	}
}

func (m *Money) Equals(val float64) bool {
	return math.Abs(m.amount.val-val) <= float64EqualityThreshold
}

func NewMoney(amount float64) *Money {
	return &Money{
		amount: &Amount{val: amount},
	}
}

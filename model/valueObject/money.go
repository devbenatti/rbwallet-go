package valueObject

import "math"

const float64EqualityThreshold = 1e-9

type Amount struct {
	val float64
}

type Money struct {
	amount Amount
}

func (m *Money) Val() float64 {
	return m.amount.val
}

func (m *Money) Add(om *Money) *Money {
	return &Money{
		amount: Amount{val: m.amount.val + om.amount.val},
	}
}

func (m *Money) Subtract(om *Money) *Money {
	return &Money{
		amount: Amount{val: m.amount.val - om.amount.val},
	}
}

func (m *Money) Negative() {
	if m.amount.val > float64(0) {
		m.amount.val = -m.amount.val
	}
}

func (m *Money) Multiply(mul int) *Money {
	return &Money{
		amount: Amount{val: m.amount.val * float64(mul)},
	}
}

func (m *Money) Equals(val float64) bool {
	return math.Abs(m.amount.val-val) <= float64EqualityThreshold
}

func NewMoney(amount float64) *Money {
	return &Money{
		amount: Amount{val: amount},
	}
}

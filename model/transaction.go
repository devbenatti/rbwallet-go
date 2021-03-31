package model

type Transaction struct {
	operation Operation
	total     *Money
}

func NewTransaction(opt int, total float64) Transaction {
	op := NewOperation(opt)
	am := NewMoney(total)

	if op.FlowType().isOutFlow() {
		am = NewMoney(-total)
	}

	return Transaction{
		operation: op,
		total:     am,
	}
}

func (t *Transaction) Operation() *Operation {
	return &t.operation
}

func (t *Transaction) Total() *Money {
	return t.total
}

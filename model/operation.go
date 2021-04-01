package model

type OperationType int
type Flow int

const (
	outFlow Flow = 1
	inFlow  Flow = 2

	cashPurchase        OperationType = 1
	installmentPurchase OperationType = 2
	withdraw            OperationType = 3
	payment             OperationType = 4
)

type Operation struct {
	operationType OperationType
	flow          Flow
}

func (fp Flow) IsOutFlow() bool {
	return fp == outFlow
}

func NewOperation(opt int) Operation {
	return Operation{
		operationType: OperationType(opt),
		flow:          getFlow(OperationType(opt)),
	}
}

func (o *Operation) Flow() Flow {
	return o.flow
}

func getFlow(ot OperationType) Flow {
	flows := map[OperationType]Flow{
		cashPurchase:        outFlow,
		installmentPurchase: outFlow,
		withdraw:            outFlow,
		payment:             inFlow,
	}

	return flows[ot]
}

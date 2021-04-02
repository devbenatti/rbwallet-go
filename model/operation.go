package model

type OperationType int
type flow int

const (
	outFlow flow = 1
	inFlow  flow = 2

	cashPurchase        OperationType = 1
	installmentPurchase OperationType = 2
	withdraw            OperationType = 3
	payment             OperationType = 4
)

func (fp flow) IsOutFlow() bool {
	return fp == outFlow
}

type Operation struct {
	operationType OperationType
	flow          flow
}

func NewOperation(opt int) Operation {
	return Operation{
		operationType: OperationType(opt),
		flow:          getFlow(OperationType(opt)),
	}
}

func (o *Operation) Flow() flow {
	return o.flow
}

func getFlow(ot OperationType) flow {
	flows := map[OperationType]flow{
		cashPurchase:        outFlow,
		installmentPurchase: outFlow,
		withdraw:            outFlow,
		payment:             inFlow,
	}

	return flows[ot]
}

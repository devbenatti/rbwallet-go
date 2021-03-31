package model

type OperationType int
type FlowType int

const (
	outFlow FlowType = 1
	inFlow  FlowType = 2

	cashPurchase        OperationType = 1
	installmentPurchase OperationType = 2
	withdraw            OperationType = 3
	payment             OperationType = 4
)

type Operation struct {
	operationType OperationType
}

func (fp FlowType) isOutFlow() bool {
	return fp == outFlow
}

func (o *Operation) Type() OperationType {
	return o.operationType
}

func NewOperation(opt int) Operation {
	return Operation{
		operationType: OperationType(opt),
	}
}

func (o *Operation) FlowType() FlowType {
	flows := getFlows()

	return flows[o.operationType]
}

func getFlows() map[OperationType]FlowType {
	flows := map[OperationType]FlowType{
		cashPurchase:        outFlow,
		installmentPurchase: outFlow,
		withdraw:            outFlow,
		payment:             inFlow,
	}

	return flows
}

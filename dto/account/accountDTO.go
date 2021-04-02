package dto

import "github.com/devbenatti/rbwallet-go/model/valueObject"

type AccountDTO struct {
	id                 valueObject.Uuid
	documentIdentifier string
}

func NewAccountDTO(id string, di string) AccountDTO {
	return AccountDTO{
		id:                 valueObject.NewUuid(id),
		documentIdentifier: di,
	}
}

func (a *AccountDTO) ID() valueObject.Uuid {
	return a.id
}

func (a *AccountDTO) DocumentIdentifier() string {
	return a.documentIdentifier
}

// asdasd

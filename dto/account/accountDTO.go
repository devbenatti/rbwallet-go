package dto

type AccountDTO struct {
	ID                 string `json:"id"`
	DocumentIdentifier string `json:"documentIdentifier"`
}

func NewAccountDTO(id string, di string) AccountDTO {
	return AccountDTO{
		ID:                 id,
		DocumentIdentifier: di,
	}
}

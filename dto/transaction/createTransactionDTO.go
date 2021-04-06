package dto

type CreateTransactionDTO struct {
	Code      string  `json:"id"`
	AccountID string  `json:"accountId" validate:"required,numeric"`
	Operation int     `json:"operationType" validate:"required,numeric,oneof=1 2 3 4"`
	Total     float64 `json:"total" validate:"required,numeric,min=0.1"`
}

package dto

type CreateTransactionDTO struct {
	Code      string  `json:"id"`
	AccountID string  `json:"accountId"`
	Operation int     `json:"operationType"`
	Total     float64 `json:"total"`
}

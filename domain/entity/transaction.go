package entity

import "time"

type Transaction struct {
	Reference       string    `json:"reference"`
	Phone           string    `json:"phone"`
	Amount          string    `json:"amount"`
	Currency        string    `json:"currency"`
	TransactionType string    `json:"transaction_type"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"update_at"`
}

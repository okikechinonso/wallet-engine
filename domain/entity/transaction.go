package entity

import "time"

type Transaction struct {
	Reference              string    `json:"reference"`
	SenderWalletAddress    string    `json:"sender_wallet_address"`
	RecipientWalletAddress string    `json:"Recipient_walllet_address"`
	UserID                 string    `json:"user_id" gorm:"foreignkey:User(id)"`
	Phone                  string    `json:"phone"`
	Amount                 int64     `json:"amount"`
	Currency               string    `json:"currency"`
	TransactionType        string    `json:"transaction_type"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"update_at"`
}

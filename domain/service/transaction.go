package service

import (
	"time"
	"wallet-engine/domain/entity"
)

func NewTransaction(reference, senderAddress, recipientAddress, userID, transactionType string, amount int64) *entity.Transaction {
	transaction := &entity.Transaction{}
	transaction.Reference = reference
	transaction.SenderWalletAddress = senderAddress
	transaction.RecipientWalletAddress = recipientAddress
	transaction.UserID = userID
	transaction.TransactionType = transactionType
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()
	transaction.Currency = "NGN"
	transaction.Amount = amount
	return transaction
}

package db

import "wallet-engine/domain/entity"

type Persistor interface {
	NewWallet(user entity.User) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	CreditdWallet(amount int64) error
	TokenInBlacklist(*string) bool
	DebitWallet(amount int64) error
	FindWallet(walletID string) (*entity.User,error)
}

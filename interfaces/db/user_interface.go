package db

import (
	"time"
	"wallet-engine/domain/entity"
)

type Persistor interface {
	NewUser(user entity.User) (*entity.User, error)
	NewWallet(walletAddress, userID string) (*entity.Wallet, error)
	FindUserByEmail(email string) (*entity.User, error)
	CreditdWallet(amount int64) error
	TokenInBlacklist(*string) bool
	DebitWallet(amount int64) error
	FindWallet(walletID string) (*entity.Wallet, error)
	GetPhone(userPhone string) (string, error)
	CreateTransaction(transaction *entity.Transaction) error
	UpdateWallet(wallet entity.Wallet) error
	ActiveWallet(wallet entity.Wallet) error
	GetMovies(page int) ([]entity.Movie, error)
	GetCommentByEmail(email string) (entity.Movie, error)
	GetCommentByDate(from, to time.Time,page int) ([]entity.Movie, error)
}

type Getter interface {
	Get(str string) error
}

package mysql

import (
	"wallet-engine/domain/entity"
)

func (d *Database) NewUser(user entity.User) (*entity.User, error) {
	result := d.PgDB.Create(&user)
	return &user, result.Error
}

func (d *Database) NewWallet(walletAddress, userID string) (*entity.Wallet, error) {
	wallet := entity.Wallet{}
	wallet.WalletAddress = walletAddress
	wallet.UserID = userID
	result := d.PgDB.Create(&wallet)
	return &wallet, result.Error
}

func (d *Database) FindWallet(walletID string) (*entity.Wallet, error) {
	var wallet entity.Wallet
	err := d.PgDB.Where("wallet_address= ?", walletID).First(&wallet).Error
	return &wallet, err
}
func (d *Database) CreditdWallet(amount int64) error {
	return nil
}
func (d *Database) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := d.PgDB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (d *Database) TokenInBlacklist(*string) bool {
	return false
}

func (d *Database) DebitWallet(amount int64) error {
	return nil
}

func (d *Database) GetPhone(userphone string) (string, error) {
	return "", nil
}

func (d *Database) CreateTransaction(transaction *entity.Transaction) error {
	result := d.PgDB.Create(transaction)
	return result.Error
}

func (d *Database) UpdateWallet(wallet entity.Wallet) error {
	result := d.PgDB.Model(&wallet).Where("wallet_address = ?", wallet.WalletAddress).Update("balance", wallet.Balance)
	return result.Error
}
func (d *Database) ActiveWallet (wallet entity.Wallet) error{
	result := d.PgDB.Model(&wallet).Where("wallet_address = ?", wallet.WalletAddress).Update("active", wallet.Active)
	return result.Error
}

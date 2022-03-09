package mysql

import "wallet-engine/domain/entity"

func (d *Database) NewWallet(wallet entity.Wallet) (*entity.Wallet,error) {
	result := d.PgDB.Create(wallet)
	return nil,result.Error
}
func (d *Database) FindWallet(walletID string) (*entity.Wallet, error) {
	var wallet entity.Wallet
	err := d.PgDB.Where("email = ?", walletID).First(&wallet).Error
	return &wallet, err
}
func ( d *Database) CreditdWallet(amount int64) error {
	return nil
}
func (d *Database) FindUserByEmail(email string) (*entity.Wallet, error){
	return nil, nil
}

func (d *Database) TokenInBlacklist(*string) bool{
return false
}

func (d *Database) DebitWallet(amount int64) error{
	return nil
}



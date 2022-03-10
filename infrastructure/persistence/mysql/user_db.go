package mysql

import "wallet-engine/domain/entity"

func (d *Database) NewWallet(user entity.User) (*entity.User,error) {
	result := d.PgDB.Create(user)
	return nil,result.Error
}
func (d *Database) FindWallet(walletID string) (*entity.User, error) {
	var user entity.W
	err := d.PgDB.Where("email = ?", walletID).First(&user).Error
	return &user, err
}
func ( d *Database) CreditdWallet(amount int64) error {
	return nil
}
func (d *Database) FindUserByEmail(email string) (*entity.User, error){
	return nil, nil
}

func (d *Database) TokenInBlacklist(*string) bool{
return false
}

func (d *Database) DebitWallet(amount int64) error{
	return nil
}



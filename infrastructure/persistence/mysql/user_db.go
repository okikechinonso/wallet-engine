package mysql

import "wallet-engine/domain/entity"

func (d *Database) NewWallet(user *entity.User) error {
	result := d.PgDB.Create(user)
	return result.Error
}
func (d *Database) FindWallet(email string) (*entity.User, error) {
	var user entity.User
	err := d.PgDB.Where("email = ?", email).First(&user).Error
	return &user, err
}

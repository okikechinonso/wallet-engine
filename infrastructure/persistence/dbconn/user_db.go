package dbconn

import (
	"kitchenmaniaapi/domain/entity"
)

func (d *Database) NewUser(user *entity.User) error {
	result := d.PgDB.Create(user)
	return result.Error
}
func (d *Database) FindUserByEmail(email string) (*entity.User,error) {
	var user entity.User
	err := d.PgDB.Where("email = ?", email).First(&user).Error
	return &user,err
}
func (d *Database) FindAllUsers() (user []*entity.User, err error){
	return user, nil
}
func (d *Database) CreateFollower(follower entity.Follower) error{
	result := d.PgDB.Create(follower)
	return result.Error
}
func (d *Database) GetAllFollowers(followedID string) error{
	return nil
}



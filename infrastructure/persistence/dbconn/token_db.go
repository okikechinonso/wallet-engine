package dbconn

import (
	"kitchenmaniaapi/domain/entity"
)

func (d *Database) TokenInBlacklist(token *string) bool {
	result := d.PgDB.Where("token = ?", *token).Find(&entity.Blacklist{})
	return result.Error != nil

}

func (d *Database) AddToBlackList(blacklist *entity.Blacklist) error {
	result := d.PgDB.Create(blacklist)
	return result.Error
}

package db

import "wallet-engine/domain/entity"

type DBinterface interface{
	NewWallet (user entity.User) (entity.User,error)
	
}
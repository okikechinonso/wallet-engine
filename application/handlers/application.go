package handlers

import (
	db "kitchenmaniaapi/infrastructure/persistence/dbinterface"
)

type App struct {
	DB db.DbInterface
}

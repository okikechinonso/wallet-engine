package handlers

import "wallet-engine/interfaces/db"

type App struct {
	DB db.Persistor
}

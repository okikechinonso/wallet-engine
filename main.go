package main

import (
	"log"
	"os"
	"wallet-engine/application/handlers"
	"wallet-engine/application/server"
	"wallet-engine/infrastructure/persistence/mysql"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Llongfile)
	env := os.Getenv("GIN_MODE")
	if env != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("couldn't load env vars: %v", err)
		}
	}
	database := &db.Database{}
	database.Init()
	database.Migrate()
	s := server.Server{App: handlers.App{DB: database}}
	s.Start()
}

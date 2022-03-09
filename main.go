package main

import (
	"kitchenmaniaapi/application/handlers"
	"kitchenmaniaapi/application/server"
	"kitchenmaniaapi/infrastructure/persistence/dbconn"
	"log"
	"os"

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
	db := &dbconn.Database{}
	db.Init()
	db.Migrate()
	s := server.Server{App: handlers.App{DB: db}}
	s.Start()
}

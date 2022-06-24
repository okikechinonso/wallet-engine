package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"wallet-engine/domain/entity"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	PgDB   *gorm.DB
	Client *mongo.Client
}

func (d *Database) Init() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_MODE")
	var dns string
	databaseurl := os.Getenv("DATABASE_URL")
	if databaseurl == "" {
		dns = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslmode)
	} else {
		dns = databaseurl
	}

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to connect to database postgresDB %v", err)
	}

	d.PgDB = db
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_PORT")))
	if err != nil {
		panic(err)
	}
	// Ping the primary
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	d.Client = client
	fmt.Println("Successfully connected to database.")
}
func (d *Database) Migrate() {

	err := d.PgDB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Printf("%s", err)
	}
	time.Sleep(time.Second)
	err = d.PgDB.AutoMigrate(&entity.Wallet{}, &entity.Transaction{}, &entity.Blacklist{}, &entity.Movie{})
	if err != nil {
		log.Printf("%s", err)
	}
}

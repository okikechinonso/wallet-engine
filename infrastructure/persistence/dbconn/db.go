package dbconn

import (
	"context"
	"fmt"
	"kitchenmaniaapi/domain/entity"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	PgDB        *gorm.DB
	MongoClient *mongo.Client
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
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatalf("unable to connect to database MongoDB %v", err)
	}
	d.MongoClient = client

	log.Println("connected to databases")
	d.PgDB = db
}
func (d *Database) Migrate() {
	err := d.PgDB.AutoMigrate(&entity.User{}, &entity.Blacklist{})
	if err != nil {
		log.Printf("%s", err)
	}
}

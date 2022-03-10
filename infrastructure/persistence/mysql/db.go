package mysql

import (
	"fmt"
	"log"
	"os"
	"wallet-engine/domain/entity"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type Database struct {
	PgDB *gorm.DB
}

func (d *Database) Init() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	var dsn string
	databaseurl := os.Getenv("DATABASE_URL")
	if databaseurl == "" {
		dsn = fmt.Sprintf( "%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user,password, dbName,)
	} else {
		dsn = databaseurl
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to connect to database postgresDB %v", err)
	}

	log.Println("connected to databases")
	d.PgDB = db
}
func (d *Database) Migrate() {
	
	err := d.PgDB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Printf("%s", err)
	}
}

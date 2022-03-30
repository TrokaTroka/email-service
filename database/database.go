package database

import (
	"log"
	"mail-sender/database/migrations"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {

	vdb := "host=localhost port=5432 user=postgres dbname=postgres password=root sslmode=disable"

	database, err := gorm.Open(postgres.Open(vdb), &gorm.Config{})

	if err != nil {
		log.Fatal("- Error: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Minute * 5)

	migrations.RunMigrations(db)
}

func GetDB() *gorm.DB {
	return db
}
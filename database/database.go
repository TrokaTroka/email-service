package database

import (
	"log"
	"mail-sender/config"
	"mail-sender/database/migrations"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {

	port, _ := config.GetConfig("db_port")
	user, _ := config.GetConfig("db_user")
	dbname, _ := config.GetConfig("db_name")
	password, _ := config.GetConfig("db_password")
	
	vdb := "host=localhost port=" + port + " user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=disable"

	database, err := gorm.Open(postgres.Open(vdb), &gorm.Config{})

	if err != nil {
		log.Fatal("- Error: ", err)
	}

	db = database

	dbConfig, _ := db.DB()

	dbConfig.SetMaxIdleConns(10)
	dbConfig.SetMaxOpenConns(100)
	dbConfig.SetConnMaxLifetime(time.Minute * 5)

	migrations.RunMigrations(db)
}

func GetDB() *gorm.DB {
	return db
}

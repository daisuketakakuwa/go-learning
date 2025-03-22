package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func SetupDb() {
	var host = os.Getenv("DB_HOST")
	var user = os.Getenv("DB_USER")
	var password = os.Getenv("DB_PASSWORD")
	var name = os.Getenv("DB_NAME")
	var port = os.Getenv("DB_PORT")
	var timezone = os.Getenv("OS_TIMEZONE")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=" + timezone
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("FAILED TO CONNECT TO DB")
	} else {
		log.Println("CONNECTED TO DB")
	}
}

func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println("LOAD .env file")
	}
}

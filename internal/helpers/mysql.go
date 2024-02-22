package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	godotenv.Load()

	dsn := os.Getenv("DSN")

	if dsn == "" {
		log.Panicln("DSN must be set")
	}

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Error connecting to database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

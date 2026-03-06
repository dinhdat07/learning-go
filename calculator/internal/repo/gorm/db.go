package gorm

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	godotenv.Load()
}

func Connect() (*gorm.DB, error) {
	dsn := os.Getenv("POSTGRES_DNS")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("[GORM] Connected to DB succesfully!")
	return db, nil
}

func Migrate() {

}

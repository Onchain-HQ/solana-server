package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DB struct {
	*gorm.DB
}

func New() (*DB, error) {
	dbURL := os.Getenv("DATABASE_URL")

	fmt.Println("Connecting to database...")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "solana-server.",
		}})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to database, running migrations...")
	err = db.AutoMigrate()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

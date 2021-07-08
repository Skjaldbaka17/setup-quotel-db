package database

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	Pool *gorm.ConnPool
	DB   *gorm.DB
}

func InitializeDBConnection() (*Connection, error) {
	godotenv.Load()
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		return &Connection{}, err
	}
	// Auto Migrate
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Quote{})

	return &Connection{
		DB: db,
	}, nil
}

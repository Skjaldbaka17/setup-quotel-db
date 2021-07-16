package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Connection struct {
	Pool *gorm.ConnPool
	DB   *gorm.DB
}

func InitializeDBConnection() (*Connection, error) {
	godotenv.Load()
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return &Connection{}, err
	}
	connection := &Connection{
		DB: db,
	}
	connection.GetShitReady()
	// Auto Migrate
	err = db.AutoMigrate(&Author{}, &Aod{}, &Aodice{}, &Quote{}, &Qod{}, &Qodice{}, &Topic{})
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	return connection, nil
}

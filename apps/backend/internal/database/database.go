package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/dhani/bill-tracker-backend/internal/config"
	"github.com/dhani/bill-tracker-backend/internal/models"
)

var DB *gorm.DB

func Connect(cfg *config.Config) *gorm.DB {
	var err error

	logLevel := logger.Silent
	if cfg.Env == "development" {
		logLevel = logger.Info
	}

	DB, err = gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
	return DB
}

func Migrate() error {
	log.Println("Running database migrations...")

	err := DB.AutoMigrate(
		&models.Company{},
		&models.User{},
		&models.Session{},
		&models.Vendor{},
		&models.Category{},
		&models.Bill{},
		&models.BillAttachment{},
		&models.BillActivity{},
	)

	if err != nil {
		return err
	}

	log.Println("Database migrations completed")
	return nil
}

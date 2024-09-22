package database

import (
	"fmt"
	"log"
	"net/url"

	"github.com/veise3/learning-record-app/config"
	"github.com/veise3/learning-record-app/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(config *config.Config) (*gorm.DB, error) {
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
	// 	config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	u, err := url.Parse(config.DATABASE_URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %v", err)
	}

	u.Scheme = "postgres"

	// Log the connection attempt (remove in production)
	log.Printf("Attempting to connect with DSN: %s", u.String())

	db, err := gorm.Open(postgres.Open(u.String()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	// Auto Migrate
	err = db.AutoMigrate(&domain.LearningRecord{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %v", err)
	}

	return db, nil
}

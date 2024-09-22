package database

import (
	"fmt"

	"github.com/veise3/learning-record-app/config"
	"github.com/veise3/learning-record-app/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(config *config.Config) (*gorm.DB, error) {
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
	// 	config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	dsn := config.DBURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

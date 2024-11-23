package db

import (
	"github.com/ayorinde-codes/real-time-delivery-tracking/models"
	"gorm.io/gorm"
)

// MigrateDB migrates the database schema for models
func MigrateDB(db *gorm.DB) {
	// AutoMigrate the models into the database
	db.AutoMigrate(&models.User{}, &models.Order{}, &models.Tracking{})
}

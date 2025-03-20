package database

import (
	"gorm.io/gorm"
	"test-monitoring/domain"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&domain.User{})
}

package db

import (
	"fmt"

	"github.com/iufb/goth/config"
	"github.com/iufb/goth/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", config.Envs.DBHost, config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName, config.Envs.DBPort)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

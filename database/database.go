package database

import (
	"fmt"

	"github.com/ab3llo/weight-tracker/config"
	"github.com/ab3llo/weight-tracker/internals/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database connection opened")

	DB.AutoMigrate(&model.Weight{})
	fmt.Println("Database Migrated")
}

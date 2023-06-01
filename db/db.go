package db

import (
	"github.com/Guilherme415/Client-Server-API/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	db, err := gorm.Open(sqlite.Open("clientServerApi.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.DolarCotation{})

	DB = db

	return nil
}

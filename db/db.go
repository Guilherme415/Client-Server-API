package db

import (
	"github.com/Guilherme415/Client-Server-API/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() error {
	db, err := gorm.Open(sqlite.Open("clientServerApi.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.DolarCotation{})

	return nil
}

package repository

import (
	"context"
	"time"

	"github.com/Guilherme415/Client-Server-API/models"
	"gorm.io/gorm"
)

type IDolarCotationRepository interface {
	Create(dolarCotation models.DolarCotation) error
}

type DolarCotationRepository struct {
	db *gorm.DB
}

func NewDolarCotationRepository(db *gorm.DB) IDolarCotationRepository {
	return &DolarCotationRepository{db}
}

func (d *DolarCotationRepository) Create(dolarCotation models.DolarCotation) error {
	ctx := context.Background()

	dbTimout := 10 * time.Second

	ctx, cancel := context.WithTimeout(ctx, dbTimout)
	defer cancel()

	return d.db.WithContext(ctx).Create(&dolarCotation).Error
}

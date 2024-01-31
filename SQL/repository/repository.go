package repository

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// NewRepository
func NewRepository(db *gorm.DB) RepositoryInterface {
	return &repository{
		db: db,
	}
}


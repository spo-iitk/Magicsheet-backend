package sync

import "gorm.io/gorm"

type RASRepository struct {
	db *gorm.DB
}

func NewRASrepository(db *gorm.DB) *RASRepository {
	return &RASRepository{
		db: db,
	}
}

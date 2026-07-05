package sync

import (
	"context"

	"github.com/spo-iitk/Magicsheet-backend/internal/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) UpsertRecruitmentCycle(ctx context.Context, rc *database.RecruitmentCycle) error {

	var existing database.RecruitmentCycle

	err := r.db.WithContext(ctx).First(&existing, rc.ID).Error

	if err == gorm.ErrRecordNotFound {
		return r.db.WithContext(ctx).Create(rc).Error
	}

	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Model(&existing).Updates(rc).Error
}

func (r *Repository) UpsertProforma(ctx context.Context, p *database.Proforma) error {

	var existing database.Proforma

	err := r.db.WithContext(ctx).First(&existing, p.ID).Error

	if err == gorm.ErrRecordNotFound {
		return r.db.WithContext(ctx).Create(p).Error
	}

	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Model(&existing).Updates(p).Error
}

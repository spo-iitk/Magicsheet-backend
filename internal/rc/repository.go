package rc

import (
	"context"

	"github.com/spo-iitk/Magicsheet-backend/internal/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetActiveRecruitmentCycles(ctx context.Context) ([]database.RecruitmentCycle, error) {
	var rcs []database.RecruitmentCycle
	err := r.db.WithContext(ctx).Where("is_active = ?", true).Find(&rcs).Error
	if err != nil {
		return nil, err
	}
	return rcs, nil
}

func (r *Repository) GetProformasByRecruitmentCycleID(ctx context.Context, rcID string) ([]database.Proforma, error) {
	var proformas []database.Proforma
	err := r.db.WithContext(ctx).Where("recruitment_cycle_id = ?", rcID).Find(&proformas).Error
	if err != nil {
		return nil, err
	}
	return proformas, nil
}

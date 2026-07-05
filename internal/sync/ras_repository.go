package sync

import (
	"context"

	"gorm.io/gorm"
)

type RASRepository struct {
	rcDB          *gorm.DB
	applicationDB *gorm.DB
	studentDB     *gorm.DB
}

func NewRASrepository(rcDB *gorm.DB, applicationDB *gorm.DB, studentDB *gorm.DB) *RASRepository {
	return &RASRepository{
		rcDB:          rcDB,
		applicationDB: applicationDB,
		studentDB:     studentDB,
	}
}

func (r *RASRepository) GetActiveRecruitmentCycles(ctx context.Context) ([]RASRecruitmentCycle, error) {
	var rcs []RASRecruitmentCycle

	err := r.rcDB.WithContext(ctx).Where("is_active = ?", true).Find(&rcs).Error

	if err != nil {
		return nil, err
	}

	return rcs, nil
}

func (r *RASRepository) GetProformas(ctx context.Context, rcID uint) {
	// uses r.applicationDB
}

func (r *RASRepository) GetStudents(ctx context.Context) {
	// uses r.studentDB
}

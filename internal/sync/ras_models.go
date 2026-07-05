package sync

import "time"

type RASRecruitmentCycle struct {
	ID           uint   `gorm:"column:id"`
	AcademicYear string `gorm:"column:academic_year"`
	Type         string `gorm:"column:type"`
	Phase        string `gorm:"column:phase"`
	IsActive     bool   `gorm:"column:is_active"`
}

type RASProforma struct {
	ID        uint      `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	RecruitmentCycleID uint   `gorm:"column:recruitment_cycle_id"`
	CompanyID          uint   `gorm:"column:company_id"`
	CompanyName        string `gorm:"column:company_name"`

	IsApproved bool `gorm:"column:is_approved"`

	Role    string `gorm:"column:role"`
	Profile string `gorm:"column:profile"`
}

func (RASRecruitmentCycle) TableName() string {
	return "recruitment_cycles"
}

func (RASProforma) TableName() string {
	return "proformas"
}

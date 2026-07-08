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
	CompanyName        string `gorm:"column:company_name"`

	IsApproved bool `gorm:"column:is_approved"`

	Role    string `gorm:"column:role"`
	Profile string `gorm:"column:profile"`
}

type RASStudent struct {
	ID                  uint   `gorm:"column:id"`
	RollNumber          string `gorm:"column:roll_no"`
	Name                string `gorm:"column:name"`
	Email               string `gorm:"column:iitk_email"`
	Phone               string `gorm:"column:phone"`
	ProgramDepartmentID uint   `gorm:"column:program_department_id"`
}

func (RASStudent) TableName() string {
	return "students"
}

func (RASRecruitmentCycle) TableName() string {
	return "recruitment_cycles"
}

func (RASProforma) TableName() string {
	return "proformas"
}

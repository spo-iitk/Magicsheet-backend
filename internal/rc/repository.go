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

// GetProformasByRole returns proformas in an RC assigned to userID with the given role.
func (r *Repository) GetProformasByRole(
	ctx context.Context,
	rcID string,
	userID uint,
	assignmentRole database.AssignmentRole,
) ([]ProformaWithAssignment, error) {
	type row struct {
		ProformaID         uint                    `gorm:"column:proforma_id"`
		Title              string                  `gorm:"column:title"`
		RoleOffered        string                  `gorm:"column:role_offered"`
		Description        string                  `gorm:"column:description"`
		ProformaType       string                  `gorm:"column:proforma_type"`
		IsInterviewActive  bool                    `gorm:"column:is_interview_active"`
		RecruitmentCycleID uint                    `gorm:"column:recruitment_cycle_id"`
		AssignmentID       uint                    `gorm:"column:assignment_id"`
		AssignmentRole     database.AssignmentRole `gorm:"column:assignment_role"`
		AssignmentIsActive bool                    `gorm:"column:assignment_is_active"`
	}

	var rows []row
	err := r.db.WithContext(ctx).
		Table("proformas").
		Select(
			"proformas.id AS proforma_id, proformas.title, proformas.role_offered, "+
				"proformas.description, proformas.proforma_type, proformas.is_interview_active, "+
				"proformas.recruitment_cycle_id, proformas.company_id, "+
				"coordinator_assignments.id AS assignment_id, "+
				"coordinator_assignments.role AS assignment_role, "+
				"coordinator_assignments.is_active AS assignment_is_active",
		).
		Joins(
			"INNER JOIN coordinator_assignments ON coordinator_assignments.proforma_id = proformas.id "+
				"AND coordinator_assignments.user_id = ? "+
				"AND coordinator_assignments.role = ?",
			userID, string(assignmentRole),
		).
		Where("proformas.recruitment_cycle_id = ?", rcID).
		Where("proformas.deleted_at IS NULL").
		Scan(&rows).Error

	if err != nil {
		return nil, err
	}

	result := make([]ProformaWithAssignment, 0, len(rows))
	for _, r := range rows {
		result = append(result, ProformaWithAssignment{
			ID:                 r.ProformaID,
			RecruitmentCycleID: r.RecruitmentCycleID,
			Title:              r.Title,
			RoleOffered:        r.RoleOffered,
			Description:        r.Description,
			ProformaType:       r.ProformaType,
			IsInterviewActive:  r.IsInterviewActive,
			AssignmentID:       r.AssignmentID,
			AssignmentRole:     string(r.AssignmentRole),
			AssignmentIsActive: r.AssignmentIsActive,
		})
	}
	return result, nil
}

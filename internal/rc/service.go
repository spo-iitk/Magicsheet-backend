package rc

import (
	"context"
	"errors"

	"github.com/spo-iitk/Magicsheet-backend/internal/database"
	"gorm.io/gorm"
)

var (
	ErrInvalidRole       = errors.New("role must be 'apc' or 'coco'")
	ErrForbiddenAssign   = errors.New("opc can assign apc or coco; apc can only assign coco")
	ErrTargetUserInvalid = errors.New("target user not found or does not have the required role")
	ErrProformaNotInRC   = errors.New("proforma does not belong to this recruitment cycle")
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListActive(ctx context.Context) ([]map[string]interface{}, error) {
	rcs, err := s.repo.GetActiveRecruitmentCycles(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]map[string]interface{}, 0, len(rcs))
	for _, rc := range rcs {
		out = append(out, map[string]interface{}{
			"id":            rc.ID,
			"name":          rc.Type + " " + rc.Phase,
			"academic_year": rc.AcademicYear,
			"type":          rc.Type,
			"phase":         rc.Phase,
			"start_date":    rc.CreatedAt.Format("02/01/2006"),
			"is_active":     rc.IsActive,
		})
	}

	return out, nil
}

// GetProformasByRole returns proformas in the given RC assigned to userID with the given role.
func (s *Service) GetProformasByRole(
	ctx context.Context,
	rcID string,
	userID uint,
	role string,
) ([]ProformaWithAssignment, error) {
	var assignmentRole database.AssignmentRole
	switch role {
	case "apc":
		assignmentRole = database.AssignmentRoleAPC
	case "coco":
		assignmentRole = database.AssignmentRoleCoCo
	default:
		return nil, ErrInvalidRole
	}
	return s.repo.GetProformasByRole(ctx, rcID, userID, assignmentRole)
}

// AssignMagicsheet assigns a proforma to a target user.
// Permission matrix: opc → apc or coco; apc → coco only.
// Upserts: reactivates an existing assignment or creates a new one.
func (s *Service) AssignMagicsheet(
	ctx context.Context,
	rcID uint,
	callerID uint,
	callerRole string,
	req AssignMagicsheetRequest,
) (*AssignMagicsheetResponse, error) {
	var assignRole database.AssignmentRole
	switch req.TargetRole {
	case "apc":
		assignRole = database.AssignmentRoleAPC
	case "coco":
		assignRole = database.AssignmentRoleCoCo
	default:
		return nil, ErrInvalidRole
	}

	switch callerRole {
	case "opc":
		// opc can assign apc or coco
	case "god":
		// god can assign apc or coco
	case "apc":
		if req.TargetRole != "coco" {
			return nil, ErrForbiddenAssign
		}
	default:
		return nil, ErrForbiddenAssign
	}

	if _, err := s.repo.GetProformaByIDAndCycle(ctx, req.ProformaID, rcID); err != nil {
		return nil, ErrProformaNotInRC
	}

	// Look up target user by email + role
	targetUser, err := s.repo.GetUserByEmailAndRole(ctx, req.TargetEmail, database.UserRole(req.TargetRole))
	if err != nil {
		return nil, ErrTargetUserInvalid
	}

	existing, err := s.repo.GetAssignment(ctx, req.ProformaID, targetUser.ID, assignRole)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existing != nil {
		existing.IsActive = true
		existing.AssignedByID = callerID
		if err := s.repo.UpdateAssignment(ctx, existing); err != nil {
			return nil, err
		}
		return toAssignResponse(existing), nil
	}

	a := &database.CoordinatorAssignment{
		ProformaID:   req.ProformaID,
		UserID:       targetUser.ID,
		Role:         assignRole,
		AssignedByID: callerID,
		IsActive:     true,
	}
	if err := s.repo.CreateAssignment(ctx, a); err != nil {
		return nil, err
	}
	return toAssignResponse(a), nil
}

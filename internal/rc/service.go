package rc

import (
	"context"
	"errors"

	"github.com/spo-iitk/Magicsheet-backend/internal/database"
)

// ErrInvalidRole is returned when the caller passes a role other than "apc" or "coco".
var ErrInvalidRole = errors.New("role must be 'apc' or 'coco'")

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

	// convert to lightweight map for JSON response
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

// GetProformasByRole returns all proformas in the given recruitment cycle that are
// assigned to userID with the specified role. Only "apc" and "coco" are accepted.
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

package rc

import (
	"context"
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

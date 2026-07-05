package sync

import (
	"context"
	"fmt"

	"github.com/spo-iitk/Magicsheet-backend/internal/database"
)

type Service struct {
	repo    *Repository
	rasRepo *RASRepository
}

func NewService(repo *Repository, rasRepo *RASRepository) *Service {
	return &Service{
		repo:    repo,
		rasRepo: rasRepo,
	}
}

func (s *Service) SyncProformas(ctx context.Context) error {
	return nil
}

func (s *Service) SyncStudents(ctx context.Context) error {
	rcs, err := s.rasRepo.GetActiveRecruitmentCycles(ctx)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", rcs)

	return nil
}

func mapRecruitmentCycle(rc RASRecruitmentCycle) database.RecruitmentCycle {
	return database.RecruitmentCycle{
		ID:           rc.ID,
		AcademicYear: rc.AcademicYear,
		Type:         rc.Type,
		Phase:        rc.Phase,
		IsActive:     rc.IsActive,
	}
}

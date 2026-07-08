package sync

import (
	"context"

	"time"

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
	rcs, err := s.rasRepo.GetActiveRecruitmentCycles(ctx)

	if err != nil {
		return err
	}

	for _, rasRc := range rcs {
		pibsRc := mapRecruitmentCycle(rasRc)

		if err := s.repo.UpsertRecruitmentCycle(ctx, &pibsRc); err != nil {
			return err
		}

		proformas, err := s.rasRepo.GetProformas(ctx, rasRc.ID)
		if err != nil {
			return err
		}

		for _, rasProformas := range proformas {
			pibsProforma := mapProforma(rasProformas)

			if err := s.repo.UpsertProforma(ctx, &pibsProforma); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Service) SyncStudents(ctx context.Context) error {

	students, err := s.rasRepo.GetStudents(ctx)

	if err != nil {
		return nil
	}

	for _, rasStudent := range students {
		pibsStudent := mapStudent(rasStudent)

		if err := s.repo.UpsertStudent(ctx, &pibsStudent); err != nil {
			return err
		}
	}
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

func mapProforma(p RASProforma) database.Proforma {
	return database.Proforma{
		ID:                 p.ID,
		RecruitmentCycleID: p.RecruitmentCycleID,


		Title:       p.CompanyName,
		RoleOffered: p.Role,
		Description: p.Profile,

		ProformaType:      "",
		IsInterviewActive: false,
		LastSyncedAt:      time.Now(),
	}
}

func mapStudent(s RASStudent) database.Student {

	program, department := getProgramDepartment(s.ProgramDepartmentID)
	return database.Student{
		ID:            s.ID,
		RollNumber:    s.RollNumber,
		Name:          s.Name,
		Department:    department,
		Program:       program,
		Email:         s.Email,
		Phone:         s.Phone,
		CurrentStatus: database.StudentAvailable,
		IsFrozen:      false,
		LastSyncedAt:  time.Now(),
	}
}

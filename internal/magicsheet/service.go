package magicsheet

import (
	"context"

	"github.com/spo-iitk/Magicsheet-backend/internal/database"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) GetMagicSheet(ctx context.Context, proformaID uint) (*GetMagicSheetResponse, error) {
	proforma, err := s.repo.GetProforma(ctx, proformaID)
	if err != nil {
		return nil, err
	}

	rounds, err := s.repo.GetInterviewRounds(ctx, proformaID)
	if err != nil {
		return nil, err
	}

	if len(rounds) == 0 {
		if err := s.repo.CreateDefaultRounds(ctx, proformaID); err != nil {
			return nil, err
		}

		rounds, err = s.repo.GetInterviewRounds(ctx, proformaID)
		if err != nil {
			return nil, err
		}
	}

	candidates, err := s.repo.GetCandidates(ctx, proformaID)
	if err != nil {
		return nil, err
	}

	response := &GetMagicSheetResponse{
		Proforma:   mapProforma(*proforma),
		Rounds:     mapRounds(rounds),
		Candidates: mapCandidates(candidates),
	}

	return response, nil
}

func (s *Service) RegisterCandidate(ctx context.Context, userID uint, proformaID uint, rollNumber string) (*CandidateDTO, error) {
	student, err := s.repo.GetStudentByRollNumber(ctx, rollNumber)

	if err != nil {
		return nil, err
	}

	exists, err := s.repo.CandidateExists(ctx, proformaID, student.ID)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrCandidateAlreadyRegistered
	}

	rounds, err := s.repo.GetInterviewRounds(ctx, proformaID)
	if err != nil {
		return nil, err
	}

	if len(rounds) == 0 {
		if err := s.repo.CreateDefaultRounds(ctx, proformaID); err != nil {
			return nil, err
		}

		rounds, err = s.repo.GetInterviewRounds(ctx, proformaID)
		if err != nil {
			return nil, err
		}
	}

	candidate, err := s.repo.RegisterCandidate(
		ctx,
		proformaID,
		student.ID,
		rounds[0].ID,
		userID,
	)

	if err != nil {
		return nil, err
	}

	dto := mapCandidate(*candidate)

	return &dto, nil

}

func (s *Service) CheckIn(
	ctx context.Context,
	proformaID uint,
	sessionID uint,
) (*InterviewSessionDTO, error) {
	panic("not implemented")
}

func (s *Service) CheckOut(
	ctx context.Context,
	proformaID uint,
	sessionID uint,
) (*InterviewSessionDTO, error) {
	panic("not implemented")
}

func (s *Service) UpdateSessionResult(
	ctx context.Context,
	proformaID uint,
	sessionID uint,
	status database.SessionStatus,
) (*InterviewSessionDTO, error) {
	panic("not implemented")
}

func (s *Service) CreateRound(
	ctx context.Context,
	proformaID uint,
	name string,
) (*RoundDTO, error) {
	panic("not implemented")
}

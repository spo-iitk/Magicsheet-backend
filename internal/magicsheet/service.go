package magicsheet

import "github.com/spo-iitk/Magicsheet-backend/internal/database"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetMagicSheet(proformaID uint) (*GetMagicSheetResponse, error) {
	panic("not implemented")
}

func (s *Service) RegisterCandidate(proformaID uint, rollNumber string) (*CandidateDTO, error) {
	panic("not implemented")
}

func (s *Service) CheckIn(proformaID uint, sessionID uint) (*InterviewSessionDTO, error) {
	panic("not implemented")
}

func (s *Service) CheckOut(proformaID uint, sessionID uint) (*InterviewSessionDTO, error) {
	panic("not implemented")
}

func (s *Service) UpdateSessionResult(
	proformaID uint,
	sessionID uint,
	status database.SessionStatus,
) (*InterviewSessionDTO, error) {
	panic("not implemented")
}

func (s *Service) CreateRound(proformaID uint, name string) (*RoundDTO, error) {
	panic("not implemented")
}

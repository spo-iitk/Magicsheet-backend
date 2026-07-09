package magicsheet

import (
	"github.com/spo-iitk/Magicsheet-backend/internal/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Authorization
func (r *Repository) HasProformaAccess(userID uint, proformaID uint) (bool, error) {
	panic("not implemented")
}

// Magic Sheet
func (r *Repository) GetProforma(proformaID uint) (*database.Proforma, error) {
	panic("not implemented")
}

func (r *Repository) GetInterviewRounds(proformaID uint) ([]database.InterviewRound, error) {
	panic("not implemented")
}

func (r *Repository) GetCandidates(proformaID uint) ([]database.ProformaCandidate, error) {
	panic("not implemented")
}

// Candidate Registration
func (r *Repository) GetStudentByRollNumber(rollNumber string) (*database.Student, error) {
	panic("not implemented")
}

func (r *Repository) CreateCandidate(candidate *database.ProformaCandidate) error {
	panic("not implemented")
}

// Interview Rounds
func (r *Repository) CreateRound(round *database.InterviewRound) error {
	panic("not implemented")
}

// Interview Sessions
func (r *Repository) GetInterviewSession(sessionID uint) (*database.InterviewSession, error) {
	panic("not implemented")
}

func (r *Repository) CreateInterviewSession(session *database.InterviewSession) error {
	panic("not implemented")
}

func (r *Repository) UpdateInterviewSession(session *database.InterviewSession) error {
	panic("not implemented")
}

// Student
func (r *Repository) UpdateStudent(student *database.Student) error {
	panic("not implemented")
}

func (r *Repository) CandidateExists(
	proformaID uint,
	studentID uint,
) (bool, error) {
	panic("not implemented")
}

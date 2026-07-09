package magicsheet

import (
	"context"

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

// can give bugs or need updates depends on count
func (r *Repository) HasProformaAccess(ctx context.Context, userID uint, proformaID uint) (bool, error) {
	var count int64

	err := r.db.WithContext(ctx).Model(&database.CoordinatorAssignment{}).Where("user_id = ? AND proforma_id = ? AND is_active = ?",
		userID,
		proformaID,
		true).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *Repository) GetProforma(ctx context.Context, proformaID uint) (*database.Proforma, error) {
	var proforma database.Proforma

	err := r.db.WithContext(ctx).First(&proforma, proformaID).Error

	if err != nil {
		return nil, err
	}

	return &proforma, nil
}

func (r *Repository) GetInterviewRounds(ctx context.Context, proformaID uint) ([]database.InterviewRound, error) {
	var rounds []database.InterviewRound

	err := r.db.WithContext(ctx).Where("proforma_id = ?", proformaID).Order("round_number ASC").Find(&rounds).Error

	if err != nil {
		return nil, err
	}

	return rounds, nil
}

func (r *Repository) GetCandidates(ctx context.Context, proformaID uint) ([]database.ProformaCandidate, error) {
	var candidates []database.ProformaCandidate

	err := r.db.WithContext(ctx).Preload("Student").Preload("InterviewSessions").Where("proforma_id = ?", proformaID).Find(&candidates).Error

	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (r *Repository) GetStudentByRollNumber(
	ctx context.Context,
	rollNumber string,
) (*database.Student, error) {
	panic("not implemented")
}

func (r *Repository) CreateCandidate(
	ctx context.Context,
	candidate *database.ProformaCandidate,
) error {
	panic("not implemented")
}

func (r *Repository) CreateRound(
	ctx context.Context,
	round *database.InterviewRound,
) error {
	panic("not implemented")
}

func (r *Repository) GetInterviewSession(
	ctx context.Context,
	sessionID uint,
) (*database.InterviewSession, error) {
	panic("not implemented")
}

func (r *Repository) CreateInterviewSession(
	ctx context.Context,
	session *database.InterviewSession,
) error {
	panic("not implemented")
}

func (r *Repository) UpdateInterviewSession(
	ctx context.Context,
	session *database.InterviewSession,
) error {
	panic("not implemented")
}

func (r *Repository) UpdateStudent(
	ctx context.Context,
	student *database.Student,
) error {
	panic("not implemented")
}

func (r *Repository) CandidateExists(
	ctx context.Context,
	proformaID uint,
	studentID uint,
) (bool, error) {
	panic("not implemented")
}

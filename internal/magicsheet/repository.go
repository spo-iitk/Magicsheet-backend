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

	err := r.db.WithContext(ctx).Preload("Student").Preload("InterviewSessions", func(db *gorm.DB) *gorm.DB {
		return db.Order("round_id ASC")
	}).Where("proforma_id = ?", proformaID).Find(&candidates).Error

	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (r *Repository) CreateDefaultRounds(ctx context.Context, proformaID uint) error {

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		rounds := []database.InterviewRound{
			{
				ProformaID:  proformaID,
				RoundNumber: 1,
				Name:        "Round 1",
			},
			{
				ProformaID:  proformaID,
				RoundNumber: 2,
				Name:        "Round 2",
			},
			{
				ProformaID:  proformaID,
				RoundNumber: 3,
				Name:        "Round 3",
			},
		}

		if err := tx.Create(&rounds).Error; err != nil {
			return err
		}

		return nil
	})

}

func (r *Repository) RegisterCandidate(ctx context.Context, proformaID uint, studentID uint, roundID uint, addedByID uint) (*database.ProformaCandidate, error) {
	var candidate database.ProformaCandidate

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		candidate = database.ProformaCandidate{
			ProformaID: proformaID,
			StudentID:  studentID,
			Source:     database.CandidateSourceManual,
			AddedByID:  &addedByID,
		}

		if err := tx.Create(&candidate).Error; err != nil {
			return err
		}

		session := database.InterviewSession{
			ProformaID:          proformaID,
			ProformaCandidateID: candidate.ID,
			RoundID:             roundID,
			Status:              database.SessionPending,
		}

		if err := tx.Create(&session).Error; err != nil {
			return err
		}

		if err := tx.Preload("Student").Preload("InterviewSessions", func(db *gorm.DB) *gorm.DB {
			return db.Order("round_id ASC")
		}).First(&candidate, candidate.ID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &candidate, nil

}

func (r *Repository) GetStudentByRollNumber(ctx context.Context, rollNumber string) (*database.Student, error) {
	var student database.Student

	err := r.db.WithContext(ctx).Where("roll_number = ?", rollNumber).First(&student).Error

	if err != nil {
		return nil, err
	}

	return &student, nil
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

func (r *Repository) CandidateExists(ctx context.Context, proformaID uint, studentID uint) (bool, error) {
	var count int64

	err := r.db.WithContext(ctx).Model(&database.ProformaCandidate{}).Where("proforma_id = ? AND student_id = ?", proformaID, studentID).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

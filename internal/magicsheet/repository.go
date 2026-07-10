package magicsheet

import (
	"context"
	"time"

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

func (r *Repository) GetStudentByRollNumber(ctx context.Context, rollNumber string) (*database.Student, error) {
	var student database.Student

	err := r.db.WithContext(ctx).Where("roll_number = ?", rollNumber).First(&student).Error

	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *Repository) CandidateExists(ctx context.Context, proformaID uint, studentID uint) (bool, error) {
	var count int64

	err := r.db.WithContext(ctx).Model(&database.ProformaCandidate{}).Where("proforma_id = ? AND student_id = ?", proformaID, studentID).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
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

func (r *Repository) CheckIn(ctx context.Context, sessionID uint) (*database.InterviewSession, error) {
	var session database.InterviewSession

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Round").Preload("ProformaCandidate.Student").First(&session, sessionID).Error; err != nil {
			return err
		}
		if session.Status != database.SessionPending {
			return ErrInvalidSessionState
		}

		if session.Round.RoundNumber > 1 {
			var previousRound database.InterviewRound

			if err := tx.Where("proforma_id = ? AND round_number = ?", session.ProformaID, session.Round.RoundNumber-1).First(&previousRound).Error; err != nil {
				return err
			}

			var previousSession database.InterviewSession

			if err := tx.Where("proforma_candidate_id = ? AND round_id = ?", session.ProformaCandidateID, previousRound.ID).First(&previousSession).Error; err != nil {
				return err
			}

			switch previousSession.Status {

			case database.SessionPassed:
				//pretty fine

			case database.SessionCheckedOut, database.SessionResultPending:
				previousSession.Status = database.SessionPassed

				if err := tx.Save(&previousSession).Error; err != nil {
					return err
				}

			case database.SessionRejected, database.SessionAbsent:
				return ErrInvalidSessionState

			default:
				return ErrInvalidSessionState
			}
		}

		now := time.Now()

		session.Status = database.SessionCheckedIn
		session.InTime = &now

		if err := tx.Save(&session).Error; err != nil {
			return err
		}

		student := session.ProformaCandidate.Student

		student.CurrentStatus = database.StudentInterviewing

		if err := tx.Save(&student).Error; err != nil {
			return err
		}

		if err := tx.Preload("Round").Preload("ProformaCandidate.Student").First(&session, sessionID).Error; err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return &session, nil

}

func (r *Repository) CheckOut(ctx context.Context, sessionID uint) (*database.InterviewSession, error) {
	var session database.InterviewSession

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Preload("ProformaCandidate.Student").First(&session, sessionID).Error; err != nil {
			return err
		}

		if session.Status != database.SessionCheckedIn {
			return ErrInvalidSessionState
		}

		now := time.Now()

		session.Status = database.SessionCheckedOut
		session.OutTime = &now

		if err := tx.Save(&session).Error; err != nil {
			return err
		}

		student := session.ProformaCandidate.Student

		student.CurrentStatus = database.StudentAvailable

		if err := tx.Save(&student).Error; err != nil {
			return err
		}

		if err := tx.Preload("ProformaCandidate.Student").First(&session, sessionID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *Repository) UpdateSessionResult(ctx context.Context, sessionID uint, status database.SessionStatus) (*database.InterviewSession, error) {

	var session database.InterviewSession

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&session, sessionID).Error; err != nil {
			return err
		}

		if session.Status != database.SessionCheckedOut {
			return ErrInvalidSessionState
		}

		switch status {
		case database.SessionPassed, database.SessionRejected, database.SessionResultPending, database.SessionAbsent:
		//valid

		default:
			return ErrInvalidSessionState
		}

		session.Status = status

		if err := tx.Save(&session).Error; err != nil {
			return err
		}

		if err := tx.First(&session, sessionID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *Repository) CreateRound(ctx context.Context, round *database.InterviewRound) error {
	return r.db.WithContext(ctx).Create(round).Error

}

func (r *Repository) UpdateRound(ctx context.Context, roundID uint,name string,) (*database.InterviewRound, error) {
	var round database.InterviewRound

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error{
		if err := tx.First(&round, roundID).Error; err != nil {
			return err
		}

		round.Name = name 

		if err := tx.Save(&round).Error; err != nil{
			return err
		}
		
		return nil 
	})

	if err != nil {
		return nil, err
	}
	
	round &round, nil 
}
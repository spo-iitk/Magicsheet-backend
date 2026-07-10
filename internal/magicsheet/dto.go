package magicsheet

import (
	"time"

	"github.com/spo-iitk/Magicsheet-backend/internal/database"
)

type GetMagicSheetResponse struct {
	Proforma   ProformaDTO    `json:"proforma"`
	Rounds     []RoundDTO     `json:"rounds"`
	Candidates []CandidateDTO `json:"candidates"`
}

type ProformaDTO struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Role  string `json:"role"`
}

type RoundDTO struct {
	ID          uint   `json:"id"`
	RoundNumber int    `json:"round_number"`
	Name        string `json:"name"`
}

type CandidateDTO struct {
	CandidateID uint                  `json:"candidate_id"`
	Student     StudentDTO            `json:"student"`
	Sessions    []InterviewSessionDTO `json:"sessions"`
}

type StudentDTO struct {
	ID           uint                   `json:"id"`
	RollNumber   string                 `json:"roll_number"`
	Name         string                 `json:"name"`
	Department   string                 `json:"department"`
	Program      string                 `json:"program"`
	Phone        string                 `json:"phone"`
	Email        string                 `json:"email"`
	GlobalStatus database.StudentStatus `json:"global_status"`
}

type InterviewSessionDTO struct {
	ID      uint                   `json:"id"`
	RoundID uint                   `json:"round_id"`
	Status  database.SessionStatus `json:"status"`

	InTime  *time.Time `json:"in_time"`
	OutTime *time.Time `json:"out_time"`

	Remarks string `json:"remarks"`
}

type RegisterCandidateRequest struct {
	RollNumber string `json:"roll_number" binding:"required"`
}

type UpdateSessionResultRequest struct {
	Status database.SessionStatus `json:"status" binding:"required"`
}

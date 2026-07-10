package magicsheet

import (
	"github.com/spo-iitk/Magicsheet-backend/internal/database"
)

func mapProforma(proforma database.Proforma) ProformaDTO {
	return ProformaDTO{
		ID:    proforma.ID,
		Title: proforma.Title,
		Role:  proforma.RoleOffered,
	}
}

func mapRounds(rounds []database.InterviewRound) []RoundDTO {
	result := make([]RoundDTO, 0, len(rounds))

	for _, round := range rounds {
		result = append(result, RoundDTO{
			ID:          round.ID,
			RoundNumber: round.RoundNumber,
			Name:        round.Name,
		})
	}

	return result
}

func mapCandidates(candidates []database.ProformaCandidate) []CandidateDTO {
	result := make([]CandidateDTO, 0, len(candidates))

	for _, candidate := range candidates {
		result = append(result, mapCandidate(candidate))
	}

	return result
}

func mapCandidate(candidate database.ProformaCandidate) CandidateDTO {
	return CandidateDTO{
		CandidateID: candidate.ID,
		Student:     mapStudent(candidate.Student),
		Sessions:    mapInterviewSessions(candidate.InterviewSessions),
	}
}

func mapStudent(student database.Student) StudentDTO {
	return StudentDTO{
		ID:           student.ID,
		RollNumber:   student.RollNumber,
		Name:         student.Name,
		Department:   student.Department,
		Program:      student.Program,
		Phone:        student.Phone,
		Email:        student.Email,
		GlobalStatus: student.CurrentStatus,
	}
}

func mapInterviewSessions(sessions []database.InterviewSession) []InterviewSessionDTO {
	result := make([]InterviewSessionDTO, 0, len(sessions))

	for _, session := range sessions {
		result = append(result, mapSession(session))
	}
	return result
}

func mapSession(session database.InterviewSession) InterviewSessionDTO {
	return InterviewSessionDTO{
		ID:      session.ID,
		RoundID: session.RoundID,
		Status:  session.Status,
		InTime:  session.InTime,
		OutTime: session.OutTime,
		Remarks: session.Remarks,
	}
}

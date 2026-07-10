package magicsheet

import "errors"

var (
	ErrCandidateAlreadyRegistered = errors.New("candidate already registered")
	ErrNoInterviewRounds          = errors.New("Has no new sessions")
	ErrInvalidSessionState        = errors.New("invalid session state") // Adding this line
)

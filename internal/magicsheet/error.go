package magicsheet

import "errors"

var (
	ErrCandidateAlreadyRegistered = errors.New("candidate already registered")
	ErrNoInterviewRounds          = errors.New("no interview rounds configured")
)
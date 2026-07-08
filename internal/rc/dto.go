package rc

// AssignMagicsheetRequest is the payload for assigning a proforma to a user.
type AssignMagicsheetRequest struct {
	ProformaID  uint   `json:"proforma_id" binding:"required"`
	TargetRole  string `json:"target_role" binding:"required"` // "apc" or "coco"
	TargetEmail string `json:"target_email" binding:"required,email"`
}

// AssignMagicsheetResponse is returned after a successful assignment.
type AssignMagicsheetResponse struct {
	AssignmentID uint   `json:"assignment_id"`
	ProformaID   uint   `json:"proforma_id"`
	AssignedToID uint   `json:"assigned_to_id"`
	Role         string `json:"role"`
	IsActive     bool   `json:"is_active"`
}

// UserOption is a single entry in the unassigned-users dropdown.
type UserOption struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ProformaWithAssignment is the projection returned by GetProformasByRole.
type ProformaWithAssignment struct {
	ID                 uint   `json:"id"`
	RecruitmentCycleID uint   `json:"recruitment_cycle_id"`
	Title              string `json:"title"`
	RoleOffered        string `json:"role_offered"`
	Description        string `json:"description"`
	ProformaType       string `json:"proforma_type"`
	IsInterviewActive  bool   `json:"is_interview_active"`
	AssignmentID       uint   `json:"assignment_id"`
	AssignmentRole     string `json:"assignment_role"`
	AssignmentIsActive bool   `json:"assignment_is_active"`
}

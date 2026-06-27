package database

import (
	"time"

	"gorm.io/gorm"
)

//custom types declared here
type UserRole string


const (
	RoleOpc UserRole = "opc"
	RoleApc UserRole = "apc"
	RoleCoco UserRole = "coco"
)



type StudentStatus string

const (
	StudentAvailable        StudentStatus = "available"
	StudentInInterview      StudentStatus = "in_interview"
	StudentWaitingNextRound StudentStatus = "waiting_for_next_round"
	StudentFrozen           StudentStatus = "frozen"
)

type SessionStatus string

const (
	SessionWaiting    SessionStatus = "waiting"
	SessionInProgress SessionStatus = "in_progress"
	SessionCleared    SessionStatus = "cleared"
	SessionHold       SessionStatus = "hold"
	SessionAbsent     SessionStatus = "absent"
)

type CandidateSource string

const (
	CandidateSourceSynced CandidateSource = "synced"
	CandidateSourceManual CandidateSource = "manual"
)


type AssignmentRole string

const (
	AssignmentRoleOPC  AssignmentRole = "opcs"
	AssignmentRoleAPC  AssignmentRole = "apc"
	AssignmentRoleCoCo AssignmentRole = "coco"
)

type SyncAction string

const (
	SyncActionCreated SyncAction = "created"
	SyncActionUpdated SyncAction = "updated"
	SyncActionDeleted SyncAction = "deleted"
)


type User struct {
	ID        uint           `gorm:"primaryKey"                                              json:"id"`
	CreatedAt time.Time      `                                                               json:"created_at"`
	UpdatedAt time.Time      `                                                               json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"                                                   json:"-"`

	Name     string   `gorm:"type:varchar(150);not null"                                     json:"name"`
	Email    string   `gorm:"type:varchar(255);uniqueIndex:idx_user_email;not null"          json:"email"`
	Role     UserRole `gorm:"type:varchar(20);not null;index:idx_user_role"                  json:"role"`
	IsActive bool     `gorm:"default:true;not null"                                          json:"is_active"`

	Assignments []CoordinatorAssignment `gorm:"foreignKey:UserID" json:"assignments,omitempty"`
}


type Company struct {
	ID        uint           `gorm:"primaryKey"                                              json:"id"`
	CreatedAt time.Time      `                                                               json:"created_at"`
	UpdatedAt time.Time      `                                                               json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"                                                   json:"-"`

	ExternalID   string    `gorm:"type:varchar(100);uniqueIndex:idx_company_ext;not null"     json:"external_id"`
	Name         string    `gorm:"type:varchar(255);not null"                                 json:"name"`
	Industry     string    `gorm:"type:varchar(150)"                                          json:"industry"`
	Website      string    `gorm:"type:varchar(500)"                                          json:"website"`
	LogoURL      string    `gorm:"type:varchar(500)"                                          json:"logo_url"`
	LastSyncedAt time.Time `gorm:"not null;index:idx_company_sync"                            json:"last_synced_at"`


	Proformas []Proforma `gorm:"foreignKey:CompanyID" json:"proformas,omitempty"`
}


type Proforma struct {
	ID        uint           `gorm:"primaryKey"                                              json:"id"`
	CreatedAt time.Time      `                                                               json:"created_at"`
	UpdatedAt time.Time      `                                                               json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"                                                   json:"-"`

	ExternalID        string    `gorm:"type:varchar(100);uniqueIndex:idx_proforma_ext;not null" json:"external_id"`
	CompanyID         uint      `gorm:"not null;index:idx_proforma_company"                     json:"company_id"`
	Company           Company   `gorm:"constraint:OnDelete:RESTRICT"                            json:"company,omitempty"`
	Title             string    `gorm:"type:varchar(255);not null"                              json:"title"`
	RoleOffered       string    `gorm:"type:varchar(255)"                                       json:"role_offered"`
	Description       string    `gorm:"type:text"                                               json:"description"`
	ProformaType      string    `gorm:"type:varchar(50);index:idx_proforma_type"                json:"proforma_type"`
	IsInterviewActive bool      `gorm:"default:false;not null;index:idx_proforma_active"        json:"is_interview_active"`
	LastSyncedAt      time.Time `gorm:"not null"                                                json:"last_synced_at"`

	// Reverse relations
	Candidates      []ProformaCandidate    `gorm:"foreignKey:ProformaID" json:"candidates,omitempty"`
	InterviewRounds []InterviewRound       `gorm:"foreignKey:ProformaID" json:"interview_rounds,omitempty"`
	Assignments     []CoordinatorAssignment `gorm:"foreignKey:ProformaID" json:"assignments,omitempty"`
}


type Student struct {
	ID        uint           `gorm:"primaryKey"                                              json:"id"`
	CreatedAt time.Time      `                                                               json:"created_at"`
	UpdatedAt time.Time      `                                                               json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"                                                   json:"-"`

	RollNumber    string        `gorm:"type:varchar(50);uniqueIndex:idx_student_roll;not null"  json:"roll_number"`
	Name          string        `gorm:"type:varchar(150);not null"                              json:"name"`
	Department    string        `gorm:"type:varchar(100);not null;index:idx_student_dept"       json:"department"`
	Program       string        `gorm:"type:varchar(100);not null"                              json:"program"`
	Email         string        `gorm:"type:varchar(255)"                                       json:"email"`
	Phone         string        `gorm:"type:varchar(20)"                                        json:"phone"`
	CurrentStatus StudentStatus `gorm:"type:varchar(30);default:'available';not null;index:idx_student_status" json:"current_status"`
	IsFrozen      bool          `gorm:"default:false;not null;index:idx_student_frozen"         json:"is_frozen"`
	LastSyncedAt  time.Time     `gorm:"not null"                                                json:"last_synced_at"`

	// Reverse relations
	Candidacies []ProformaCandidate `gorm:"foreignKey:StudentID" json:"candidacies,omitempty"`
	Sessions    []InterviewSession  `gorm:"foreignKey:StudentID" json:"sessions,omitempty"`
}


type ProformaCandidate struct {
	ID        uint      `gorm:"primaryKey"                                                                  json:"id"`
	CreatedAt time.Time `                                                                                   json:"created_at"`
	UpdatedAt time.Time `                                                                                   json:"updated_at"`

	ProformaID uint            `gorm:"not null;uniqueIndex:idx_candidate_proforma_student"                      json:"proforma_id"`
	Proforma   Proforma        `gorm:"constraint:OnDelete:CASCADE"                                             json:"proforma,omitempty"`
	StudentID  uint            `gorm:"not null;uniqueIndex:idx_candidate_proforma_student;index:idx_candidate_student" json:"student_id"`
	Student    Student         `gorm:"constraint:OnDelete:RESTRICT"                                            json:"student,omitempty"`
	Source     CandidateSource `gorm:"type:varchar(20);not null;index:idx_candidate_source"                     json:"source"`
	AddedByID  *uint           `gorm:"index:idx_candidate_added_by"                                            json:"added_by_id"`
	AddedBy    *User           `gorm:"constraint:OnDelete:SET NULL"                                            json:"added_by,omitempty"`
}


type InterviewRound struct {
	ID        uint      `gorm:"primaryKey"                                                                  json:"id"`
	CreatedAt time.Time `                                                                                   json:"created_at"`
	UpdatedAt time.Time `                                                                                   json:"updated_at"`

	ProformaID  uint     `gorm:"not null;uniqueIndex:idx_round_proforma_number;index:idx_round_proforma"       json:"proforma_id"`
	Proforma    Proforma `gorm:"constraint:OnDelete:CASCADE"                                                  json:"proforma,omitempty"`
	Name        string   `gorm:"type:varchar(100);not null"                                                    json:"name"`
	RoundNumber int      `gorm:"not null;uniqueIndex:idx_round_proforma_number"                                json:"round_number"`

	// Reverse relations
	Sessions []InterviewSession `gorm:"foreignKey:RoundID" json:"sessions,omitempty"`
}


type InterviewSession struct {
	ID        uint      `gorm:"primaryKey"                                                                  json:"id"`
	CreatedAt time.Time `                                                                                   json:"created_at"`
	UpdatedAt time.Time `                                                                                   json:"updated_at"`

	StudentID     uint           `gorm:"not null;index:idx_session_student;uniqueIndex:idx_session_student_round"  json:"student_id"`
	Student       Student        `gorm:"constraint:OnDelete:RESTRICT"                                             json:"student,omitempty"`
	ProformaID    uint           `gorm:"not null;index:idx_session_proforma"                                      json:"proforma_id"`
	Proforma      Proforma       `gorm:"constraint:OnDelete:RESTRICT"                                             json:"proforma,omitempty"`
	RoundID       uint           `gorm:"not null;index:idx_session_round;uniqueIndex:idx_session_student_round"   json:"round_id"`
	Round         InterviewRound `gorm:"constraint:OnDelete:RESTRICT"                                             json:"round,omitempty"`
	ConductedByID uint           `gorm:"not null;index:idx_session_conductor"                                     json:"conducted_by_id"`
	ConductedBy   User           `gorm:"constraint:OnDelete:RESTRICT"                                             json:"conducted_by,omitempty"`
	InTime        *time.Time     `gorm:"index:idx_session_intime"                                                  json:"in_time"`
	OutTime       *time.Time     `                                                                                 json:"out_time"`
	Status        SessionStatus  `gorm:"type:varchar(20);default:'waiting';not null;index:idx_session_status"      json:"status"`
	Remarks       string         `gorm:"type:text"                                                                 json:"remarks"`
}


type CoordinatorAssignment struct {
	ID        uint      `gorm:"primaryKey"                                                                  json:"id"`
	CreatedAt time.Time `                                                                                   json:"created_at"`
	UpdatedAt time.Time `                                                                                   json:"updated_at"`

	ProformaID   uint           `gorm:"not null;uniqueIndex:idx_assign_proforma_user_role;index:idx_assign_proforma" json:"proforma_id"`
	Proforma     Proforma       `gorm:"constraint:OnDelete:CASCADE"                                                json:"proforma,omitempty"`
	UserID       uint           `gorm:"not null;uniqueIndex:idx_assign_proforma_user_role;index:idx_assign_user"    json:"user_id"`
	User         User           `gorm:"constraint:OnDelete:RESTRICT"                                               json:"user,omitempty"`
	Role         AssignmentRole `gorm:"type:varchar(10);not null;uniqueIndex:idx_assign_proforma_user_role"         json:"role"`
	AssignedByID uint           `gorm:"not null;index:idx_assign_by"                                               json:"assigned_by_id"`
	AssignedBy   User           `gorm:"foreignKey:AssignedByID;constraint:OnDelete:RESTRICT"                       json:"assigned_by,omitempty"`
	IsActive     bool           `gorm:"default:true;not null"                                                      json:"is_active"`
}


type SyncLog struct {
	ID        uint      `gorm:"primaryKey"                                              json:"id"`
	CreatedAt time.Time `                                                               json:"created_at"`

	EntityType   string     `gorm:"type:varchar(50);not null;index:idx_sync_entity"         json:"entity_type"`
	ExternalID   string     `gorm:"type:varchar(100);index:idx_sync_external"               json:"external_id"`
	Action       SyncAction `gorm:"type:varchar(20);not null"                               json:"action"`
	RecordsCount int        `gorm:"default:0;not null"                                      json:"records_count"`
	Status       string     `gorm:"type:varchar(20);not null;index:idx_sync_status"         json:"status"`
	ErrorMessage string     `gorm:"type:text"                                               json:"error_message"`
	SyncDuration int        `gorm:"not null"                                                json:"sync_duration_ms"`
}


func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Company{},
		&Proforma{},
		&Student{},
		&ProformaCandidate{},
		&InterviewRound{},
		&InterviewSession{},
		&CoordinatorAssignment{},
		&SyncLog{},
	)
}


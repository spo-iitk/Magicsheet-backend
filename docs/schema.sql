-- PostgreSQL DDL Schema for MagicSheet Backend
-- Generated from GORM models (Updated Version)

-- =========================================================================
-- 1. INDEPENDENT TABLES
-- =========================================================================

-- Table: users
CREATE TABLE users (
    id bigserial PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    name varchar(150) NOT NULL,
    email varchar(255) NOT NULL,
    role varchar(20) NOT NULL,
    is_active boolean NOT NULL DEFAULT true,
    hash_password text NOT NULL
);

CREATE INDEX idx_users_deleted_at ON users (deleted_at);
CREATE UNIQUE INDEX idx_user_email ON users (email);
CREATE INDEX idx_user_role ON users (role);

-- Table: recruitment_cycles
CREATE TABLE recruitment_cycles (
    id bigserial PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    academic_year varchar(20) NOT NULL,
    type varchar(50) NOT NULL,
    phase varchar(50) NOT NULL,
    is_active boolean NOT NULL DEFAULT false
);

CREATE INDEX idx_recruitment_cycles_deleted_at ON recruitment_cycles (deleted_at);
CREATE INDEX idx_rc_active ON recruitment_cycles (is_active);

-- Table: students
CREATE TABLE students (
    id bigserial PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    roll_number varchar(50) NOT NULL,
    name varchar(150) NOT NULL,
    department varchar(100) NOT NULL,
    program varchar(100) NOT NULL,
    email varchar(255),
    phone varchar(20),
    current_status varchar(30) NOT NULL DEFAULT 'available',
    is_frozen boolean NOT NULL DEFAULT false,
    last_synced_at timestamptz NOT NULL
);

CREATE INDEX idx_students_deleted_at ON students (deleted_at);
CREATE UNIQUE INDEX idx_student_roll ON students (roll_number);
CREATE INDEX idx_student_dept ON students (department);
CREATE INDEX idx_student_status ON students (current_status);
CREATE INDEX idx_student_frozen ON students (is_frozen);

-- Table: sync_logs
CREATE TABLE sync_logs (
    id bigserial PRIMARY KEY,
    created_at timestamptz,
    entity_type varchar(50) NOT NULL,
    external_id varchar(100),
    action varchar(20) NOT NULL,
    records_count integer NOT NULL DEFAULT 0,
    status varchar(20) NOT NULL,
    error_message text,
    sync_duration integer NOT NULL
);

CREATE INDEX idx_sync_entity ON sync_logs (entity_type);
CREATE INDEX idx_sync_external ON sync_logs (external_id);
CREATE INDEX idx_sync_status ON sync_logs (status);


-- =========================================================================
-- 2. DEPENDENT TABLES (With Foreign Keys)
-- =========================================================================

-- Table: proformas
CREATE TABLE proformas (
    id bigserial PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    recruitment_cycle_id bigint NOT NULL,
    external_id varchar(225) NOT NULL,
    title varchar(255) NOT NULL,
    role_offered varchar(255),
    description text,
    proforma_type varchar(50),
    is_interview_active boolean NOT NULL DEFAULT false,
    last_synced_at timestamptz NOT NULL,
    
    CONSTRAINT fk_recruitment_cycles_proformas 
        FOREIGN KEY (recruitment_cycle_id) 
        REFERENCES recruitment_cycles (id) 
        ON DELETE RESTRICT
);

CREATE INDEX idx_proformas_deleted_at ON proformas (deleted_at);
CREATE INDEX idx_proforma_cycle ON proformas (recruitment_cycle_id);
CREATE UNIQUE INDEX idx_proforma_ext ON proformas (external_id);
CREATE INDEX idx_proforma_type ON proformas (proforma_type);
CREATE INDEX idx_proforma_active ON proformas (is_interview_active);

-- Table: proforma_candidates
CREATE TABLE proforma_candidates (
    id bigserial PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    proforma_id bigint NOT NULL,
    student_id bigint NOT NULL,
    source varchar(20) NOT NULL,
    added_by_id bigint,
    
    CONSTRAINT fk_proformas_candidates 
        FOREIGN KEY (proforma_id) 
        REFERENCES proformas (id) 
        ON DELETE CASCADE,
        
    CONSTRAINT fk_students_candidates 
        FOREIGN KEY (student_id) 
        REFERENCES students (id) 
        ON DELETE RESTRICT,
        
    CONSTRAINT fk_users_candidates 
        FOREIGN KEY (added_by_id) 
        REFERENCES users (id) 
        ON DELETE SET NULL
);

CREATE UNIQUE INDEX idx_candidate_proforma_student ON proforma_candidates (proforma_id, student_id);
CREATE INDEX idx_candidate_student ON proforma_candidates (student_id);
CREATE INDEX idx_candidate_source ON proforma_candidates (source);
CREATE INDEX idx_candidate_added_by ON proforma_candidates (added_by_id);

-- Table: interview_rounds
CREATE TABLE interview_rounds (
    id bigserial PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    proforma_id bigint NOT NULL,
    name varchar(100) NOT NULL,
    round_number integer NOT NULL,
    
    CONSTRAINT fk_proformas_interview_rounds 
        FOREIGN KEY (proforma_id) 
        REFERENCES proformas (id) 
        ON DELETE CASCADE
);

CREATE UNIQUE INDEX idx_round_proforma_number ON interview_rounds (proforma_id, round_number);
CREATE INDEX idx_round_proforma ON interview_rounds (proforma_id);

-- Table: interview_sessions
CREATE TABLE interview_sessions (
    id bigserial PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    proforma_candidate_id bigint NOT NULL,
    proforma_id bigint NOT NULL,
    round_id bigint NOT NULL,
    conducted_by_id bigint NOT NULL,
    in_time timestamptz,
    out_time timestamptz,
    status varchar(20) NOT NULL DEFAULT 'waiting',
    remarks text,
    
    CONSTRAINT fk_proforma_candidates_interview_sessions 
        FOREIGN KEY (proforma_candidate_id) 
        REFERENCES proforma_candidates (id) 
        ON DELETE RESTRICT,

    CONSTRAINT fk_proformas_interview_sessions 
        FOREIGN KEY (proforma_id) 
        REFERENCES proformas (id) 
        ON DELETE CASCADE,
        
    CONSTRAINT fk_interview_rounds_interview_sessions 
        FOREIGN KEY (round_id) 
        REFERENCES interview_rounds (id) 
        ON DELETE RESTRICT,
        
    CONSTRAINT fk_users_interview_sessions 
        FOREIGN KEY (conducted_by_id) 
        REFERENCES users (id) 
        ON DELETE RESTRICT
);

CREATE UNIQUE INDEX idx_session_candidate_round ON interview_sessions (proforma_candidate_id, round_id);
CREATE INDEX idx_session_candidate ON interview_sessions (proforma_candidate_id);
CREATE UNIQUE INDEX idx_candidate_proforma_student ON interview_sessions (proforma_id);
CREATE INDEX idx_session_round ON interview_sessions (round_id);
CREATE INDEX idx_session_conductor ON interview_sessions (conducted_by_id);
CREATE INDEX idx_session_intime ON interview_sessions (in_time);
CREATE INDEX idx_session_status ON interview_sessions (status);

-- Table: coordinator_assignments
CREATE TABLE coordinator_assignments (
    id bigserial PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    proforma_id bigint NOT NULL,
    user_id bigint NOT NULL,
    role varchar(10) NOT NULL,
    assigned_by_id bigint NOT NULL,
    is_active boolean NOT NULL DEFAULT true,
    
    CONSTRAINT fk_proformas_coordinator_assignments 
        FOREIGN KEY (proforma_id) 
        REFERENCES proformas (id) 
        ON DELETE CASCADE,
        
    CONSTRAINT fk_users_coordinator_assignments 
        FOREIGN KEY (user_id) 
        REFERENCES users (id) 
        ON DELETE RESTRICT,
        
    CONSTRAINT fk_users_assigned_by_coordinator_assignments 
        FOREIGN KEY (assigned_by_id) 
        REFERENCES users (id) 
        ON DELETE RESTRICT
);

CREATE UNIQUE INDEX idx_assign_proforma_user_role ON coordinator_assignments (proforma_id, user_id, role);
CREATE INDEX idx_assign_proforma ON coordinator_assignments (proforma_id);
CREATE INDEX idx_assign_user ON coordinator_assignments (user_id);
CREATE INDEX idx_assign_by ON coordinator_assignments (assigned_by_id);

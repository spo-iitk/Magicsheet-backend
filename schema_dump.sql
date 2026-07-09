--
-- PostgreSQL database dump
--

\restrict QFRopgHSF0O0Tb3JmsmxOJQgbXvCSLHvZHoyJSbKswgs7YtEGoAbMEb0bx4bT2R

-- Dumped from database version 17.10
-- Dumped by pg_dump version 17.10

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: admin
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO admin;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: admin
--

COMMENT ON SCHEMA public IS '';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: coordinator_assignments; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.coordinator_assignments (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    proforma_id bigint NOT NULL,
    user_id bigint NOT NULL,
    role character varying(10) NOT NULL,
    assigned_by_id bigint NOT NULL,
    is_active boolean DEFAULT true NOT NULL
);


ALTER TABLE public.coordinator_assignments OWNER TO admin;

--
-- Name: coordinator_assignments_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.coordinator_assignments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.coordinator_assignments_id_seq OWNER TO admin;

--
-- Name: coordinator_assignments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.coordinator_assignments_id_seq OWNED BY public.coordinator_assignments.id;


--
-- Name: interview_rounds; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.interview_rounds (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    proforma_id bigint NOT NULL,
    name character varying(100) NOT NULL,
    round_number bigint NOT NULL
);


ALTER TABLE public.interview_rounds OWNER TO admin;

--
-- Name: interview_rounds_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.interview_rounds_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.interview_rounds_id_seq OWNER TO admin;

--
-- Name: interview_rounds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.interview_rounds_id_seq OWNED BY public.interview_rounds.id;


--
-- Name: interview_sessions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.interview_sessions (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    proforma_candidate_id bigint NOT NULL,
    proforma_id bigint NOT NULL,
    round_id bigint NOT NULL,
    conducted_by_id bigint NOT NULL,
    in_time timestamp with time zone,
    out_time timestamp with time zone,
    status character varying(20) DEFAULT 'waiting'::character varying NOT NULL,
    remarks text
);


ALTER TABLE public.interview_sessions OWNER TO admin;

--
-- Name: interview_sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.interview_sessions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.interview_sessions_id_seq OWNER TO admin;

--
-- Name: interview_sessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.interview_sessions_id_seq OWNED BY public.interview_sessions.id;


--
-- Name: proforma_candidates; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.proforma_candidates (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    proforma_id bigint NOT NULL,
    student_id bigint NOT NULL,
    source character varying(20) NOT NULL,
    added_by_id bigint
);


ALTER TABLE public.proforma_candidates OWNER TO admin;

--
-- Name: proforma_candidates_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.proforma_candidates_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.proforma_candidates_id_seq OWNER TO admin;

--
-- Name: proforma_candidates_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.proforma_candidates_id_seq OWNED BY public.proforma_candidates.id;


--
-- Name: proformas; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.proformas (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    recruitment_cycle_id bigint NOT NULL,
    company_id bigint NOT NULL,
    title character varying(255) NOT NULL,
    role_offered character varying(255),
    description text,
    proforma_type character varying(50),
    is_interview_active boolean DEFAULT false NOT NULL,
    last_synced_at timestamp with time zone NOT NULL
);


ALTER TABLE public.proformas OWNER TO admin;

--
-- Name: proformas_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.proformas_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.proformas_id_seq OWNER TO admin;

--
-- Name: proformas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.proformas_id_seq OWNED BY public.proformas.id;


--
-- Name: recruitment_cycles; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.recruitment_cycles (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    academic_year character varying(20) NOT NULL,
    type character varying(50) NOT NULL,
    phase character varying(50) NOT NULL,
    is_active boolean DEFAULT false NOT NULL
);


ALTER TABLE public.recruitment_cycles OWNER TO admin;

--
-- Name: recruitment_cycles_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.recruitment_cycles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.recruitment_cycles_id_seq OWNER TO admin;

--
-- Name: recruitment_cycles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.recruitment_cycles_id_seq OWNED BY public.recruitment_cycles.id;


--
-- Name: students; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.students (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    roll_number character varying(50) NOT NULL,
    name character varying(150) NOT NULL,
    department character varying(100) NOT NULL,
    program character varying(100) NOT NULL,
    email character varying(255),
    phone character varying(20),
    current_status character varying(30) DEFAULT 'available'::character varying NOT NULL,
    is_frozen boolean DEFAULT false NOT NULL,
    last_synced_at timestamp with time zone NOT NULL
);


ALTER TABLE public.students OWNER TO admin;

--
-- Name: students_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.students_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.students_id_seq OWNER TO admin;

--
-- Name: students_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.students_id_seq OWNED BY public.students.id;


--
-- Name: sync_logs; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sync_logs (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    entity_type character varying(50) NOT NULL,
    external_id character varying(100),
    action character varying(20) NOT NULL,
    records_count bigint DEFAULT 0 NOT NULL,
    status character varying(20) NOT NULL,
    error_message text,
    sync_duration bigint NOT NULL
);


ALTER TABLE public.sync_logs OWNER TO admin;

--
-- Name: sync_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.sync_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sync_logs_id_seq OWNER TO admin;

--
-- Name: sync_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.sync_logs_id_seq OWNED BY public.sync_logs.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(150) NOT NULL,
    email character varying(255) NOT NULL,
    role character varying(20) NOT NULL,
    is_active boolean DEFAULT true NOT NULL,
    hash_password text NOT NULL
);


ALTER TABLE public.users OWNER TO admin;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO admin;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: coordinator_assignments id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.coordinator_assignments ALTER COLUMN id SET DEFAULT nextval('public.coordinator_assignments_id_seq'::regclass);


--
-- Name: interview_rounds id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.interview_rounds ALTER COLUMN id SET DEFAULT nextval('public.interview_rounds_id_seq'::regclass);


--
-- Name: interview_sessions id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.interview_sessions ALTER COLUMN id SET DEFAULT nextval('public.interview_sessions_id_seq'::regclass);


--
-- Name: proforma_candidates id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.proforma_candidates ALTER COLUMN id SET DEFAULT nextval('public.proforma_candidates_id_seq'::regclass);


--
-- Name: proformas id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.proformas ALTER COLUMN id SET DEFAULT nextval('public.proformas_id_seq'::regclass);


--
-- Name: recruitment_cycles id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.recruitment_cycles ALTER COLUMN id SET DEFAULT nextval('public.recruitment_cycles_id_seq'::regclass);


--
-- Name: students id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.students ALTER COLUMN id SET DEFAULT nextval('public.students_id_seq'::regclass);


--
-- Name: sync_logs id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sync_logs ALTER COLUMN id SET DEFAULT nextval('public.sync_logs_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: coordinator_assignments coordinator_assignments_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.coordinator_assignments
    ADD CONSTRAINT coordinator_assignments_pkey PRIMARY KEY (id);


--
-- Name: interview_rounds interview_rounds_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.interview_rounds
    ADD CONSTRAINT interview_rounds_pkey PRIMARY KEY (id);


--
-- Name: interview_sessions interview_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.interview_sessions
    ADD CONSTRAINT interview_sessions_pkey PRIMARY KEY (id);


--
-- Name: proforma_candidates proforma_candidates_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.proforma_candidates
    ADD CONSTRAINT proforma_candidates_pkey PRIMARY KEY (id);


--
-- Name: proformas proformas_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.proformas
    ADD CONSTRAINT proformas_pkey PRIMARY KEY (id);


--
-- Name: recruitment_cycles recruitment_cycles_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.recruitment_cycles
    ADD CONSTRAINT recruitment_cycles_pkey PRIMARY KEY (id);


--
-- Name: students students_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.students
    ADD CONSTRAINT students_pkey PRIMARY KEY (id);


--
-- Name: sync_logs sync_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sync_logs
    ADD CONSTRAINT sync_logs_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_assign_by; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_assign_by ON public.coordinator_assignments USING btree (assigned_by_id);


--
-- Name: idx_assign_proforma; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_assign_proforma ON public.coordinator_assignments USING btree (proforma_id);


--
-- Name: idx_assign_proforma_user_role; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX idx_assign_proforma_user_role ON public.coordinator_assignments USING btree (proforma_id, user_id, role);


--
-- Name: idx_assign_user; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_assign_user ON public.coordinator_assignments USING btree (user_id);


--
-- Name: idx_candidate_added_by; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_candidate_added_by ON public.proforma_candidates USING btree (added_by_id);


--
-- Name: idx_candidate_proforma_student; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX idx_candidate_proforma_student ON public.proforma_candidates USING btree (proforma_id, student_id);


--
-- Name: idx_candidate_source; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_candidate_source ON public.proforma_candidates USING btree (source);


--
-- Name: idx_candidate_student; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_candidate_student ON public.proforma_candidates USING btree (student_id);


--
-- Name: idx_proforma_active; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_proforma_active ON public.proformas USING btree (is_interview_active);


--
-- Name: idx_proforma_company; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_proforma_company ON public.proformas USING btree (company_id);


--
-- Name: idx_proforma_cycle; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_proforma_cycle ON public.proformas USING btree (recruitment_cycle_id);


--
-- Name: idx_proforma_type; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_proforma_type ON public.proformas USING btree (proforma_type);


--
-- Name: idx_proformas_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_proformas_deleted_at ON public.proformas USING btree (deleted_at);


--
-- Name: idx_rc_active; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_rc_active ON public.recruitment_cycles USING btree (is_active);


--
-- Name: idx_recruitment_cycles_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_recruitment_cycles_deleted_at ON public.recruitment_cycles USING btree (deleted_at);


--
-- Name: idx_round_proforma; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_round_proforma ON public.interview_rounds USING btree (proforma_id);


--
-- Name: idx_round_proforma_number; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX idx_round_proforma_number ON public.interview_rounds USING btree (proforma_id, round_number);


--
-- Name: idx_session_candidate; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_session_candidate ON public.interview_sessions USING btree (proforma_candidate_id);


--
-- Name: idx_session_candidate_round; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX idx_session_candidate_round ON public.interview_sessions USING btree (proforma_candidate_id, round_id);


--
-- Name: idx_session_conductor; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_session_conductor ON public.interview_sessions USING btree (conducted_by_id);


--
-- Name: idx_session_intime; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_session_intime ON public.interview_sessions USING btree (in_time);


--
-- Name: idx_session_round; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_session_round ON public.interview_sessions USING btree (round_id);


--
-- Name: idx_session_status; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_session_status ON public.interview_sessions USING btree (status);


--
-- Name: idx_student_dept; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_student_dept ON public.students USING btree (department);


--
-- Name: idx_student_frozen; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_student_frozen ON public.students USING btree (is_frozen);


--
-- Name: idx_student_roll; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX idx_student_roll ON public.students USING btree (roll_number);


--
-- Name: idx_student_status; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_student_status ON public.students USING btree (current_status);


--
-- Name: idx_students_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_students_deleted_at ON public.students USING btree (deleted_at);


--
-- Name: idx_sync_entity; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_sync_entity ON public.sync_logs USING btree (entity_type);


--
-- Name: idx_sync_external; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_sync_external ON public.sync_logs USING btree (external_id);


--
-- Name: idx_sync_status; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_sync_status ON public.sync_logs USING btree (status);


--
-- Name: idx_user_email; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX idx_user_email ON public.users USING btree (email);


--
-- Name: idx_user_role; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_user_role ON public.users USING btree (role);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: coordinator_assignments fk_coordinator_assignments_assigned_by; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.coordinator_assignments
    ADD CONSTRAINT fk_coordinator_assignments_assigned_by FOREIGN KEY (assigned_by_id) REFERENCES public.users(id) ON DELETE RESTRICT;


--
-- Name: interview_sessions fk_interview_rounds_sessions; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.interview_sessions
    ADD CONSTRAINT fk_interview_rounds_sessions FOREIGN KEY (round_id) REFERENCES public.interview_rounds(id);


--
-- Name: interview_sessions fk_interview_sessions_conducted_by; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.interview_sessions
    ADD CONSTRAINT fk_interview_sessions_conducted_by FOREIGN KEY (conducted_by_id) REFERENCES public.users(id) ON DELETE RESTRICT;


--
-- Name: interview_sessions fk_interview_sessions_proforma; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.interview_sessions
    ADD CONSTRAINT fk_interview_sessions_proforma FOREIGN KEY (proforma_id) REFERENCES public.proformas(id) ON DELETE CASCADE;


--
-- Name: proforma_candidates fk_proforma_candidates_added_by; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.proforma_candidates
    ADD CONSTRAINT fk_proforma_candidates_added_by FOREIGN KEY (added_by_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: interview_sessions fk_proforma_candidates_interview_sessions; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.interview_sessions
    ADD CONSTRAINT fk_proforma_candidates_interview_sessions FOREIGN KEY (proforma_candidate_id) REFERENCES public.proforma_candidates(id);


--
-- Name: proforma_candidates fk_proformas_candidates; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.proforma_candidates
    ADD CONSTRAINT fk_proformas_candidates FOREIGN KEY (proforma_id) REFERENCES public.proformas(id);


--
-- Name: coordinator_assignments fk_proformas_coordinator_assignments; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.coordinator_assignments
    ADD CONSTRAINT fk_proformas_coordinator_assignments FOREIGN KEY (proforma_id) REFERENCES public.proformas(id);


--
-- Name: interview_rounds fk_proformas_interview_rounds; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.interview_rounds
    ADD CONSTRAINT fk_proformas_interview_rounds FOREIGN KEY (proforma_id) REFERENCES public.proformas(id);


--
-- Name: proformas fk_recruitment_cycles_proformas; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.proformas
    ADD CONSTRAINT fk_recruitment_cycles_proformas FOREIGN KEY (recruitment_cycle_id) REFERENCES public.recruitment_cycles(id);


--
-- Name: proforma_candidates fk_students_candidacies; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.proforma_candidates
    ADD CONSTRAINT fk_students_candidacies FOREIGN KEY (student_id) REFERENCES public.students(id);


--
-- Name: coordinator_assignments fk_users_assignments; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.coordinator_assignments
    ADD CONSTRAINT fk_users_assignments FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: admin
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;


--
-- PostgreSQL database dump complete
--

\unrestrict QFRopgHSF0O0Tb3JmsmxOJQgbXvCSLHvZHoyJSbKswgs7YtEGoAbMEb0bx4bT2R


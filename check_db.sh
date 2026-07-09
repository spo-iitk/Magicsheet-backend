#!/bin/bash
DB="docker exec portal-postgres psql -U admin -d portal"

echo "=== Sample Proformas ===" > db_samples.txt
$DB -c "SELECT * FROM proformas LIMIT 5;" >> db_samples.txt
echo "=== Sample Students ===" >> db_samples.txt
$DB -c "SELECT * FROM students LIMIT 5;" >> db_samples.txt
echo "=== Sample Proforma Candidates ===" >> db_samples.txt
$DB -c "SELECT * FROM proforma_candidates LIMIT 5;" >> db_samples.txt
echo "=== Sample Interview Rounds ===" >> db_samples.txt
$DB -c "SELECT * FROM interview_rounds LIMIT 5;" >> db_samples.txt
echo "=== Sample Interview Sessions ===" >> db_samples.txt
$DB -c "SELECT * FROM interview_sessions LIMIT 5;" >> db_samples.txt
echo "=== Sample Coordinator Assignments ===" >> db_samples.txt
$DB -c "SELECT * FROM coordinator_assignments LIMIT 5;" >> db_samples.txt

echo "=== Checks ===" > db_checks.txt
echo "Check unexpected NULL values in interview_sessions.round_id or proforma_id" >> db_checks.txt
$DB -c "SELECT count(*) FROM interview_sessions WHERE round_id IS NULL OR proforma_id IS NULL;" >> db_checks.txt

echo "Check unexpected NULLs in proforma_candidates.student_id or proforma_id" >> db_checks.txt
$DB -c "SELECT count(*) FROM proforma_candidates WHERE student_id IS NULL OR proforma_id IS NULL;" >> db_checks.txt

echo "Check for duplicate interview_sessions for same candidate and round" >> db_checks.txt
$DB -c "SELECT proforma_candidate_id, round_id, count(*) FROM interview_sessions GROUP BY proforma_candidate_id, round_id HAVING count(*) > 1;" >> db_checks.txt

echo "Check for inconsistent proforma_id in interview_sessions (mismatch with proforma_candidates)" >> db_checks.txt
$DB -c "SELECT s.id, s.proforma_id AS session_proforma, c.proforma_id AS candidate_proforma FROM interview_sessions s JOIN proforma_candidates c ON s.proforma_candidate_id = c.id WHERE s.proforma_id != c.proforma_id;" >> db_checks.txt

echo "Check for inconsistent proforma_id in interview_sessions (mismatch with interview_rounds)" >> db_checks.txt
$DB -c "SELECT s.id, s.proforma_id AS session_proforma, r.proforma_id AS round_proforma FROM interview_sessions s JOIN interview_rounds r ON s.round_id = r.id WHERE s.proforma_id != r.proforma_id;" >> db_checks.txt

echo "Check for duplicate rounds per proforma" >> db_checks.txt
$DB -c "SELECT proforma_id, round_number, count(*) FROM interview_rounds GROUP BY proforma_id, round_number HAVING count(*) > 1;" >> db_checks.txt

echo "Check for orphaned candidates (student deleted)" >> db_checks.txt
$DB -c "SELECT c.id, c.student_id FROM proforma_candidates c LEFT JOIN students s ON c.student_id = s.id WHERE s.id IS NULL;" >> db_checks.txt

echo "Check missing user for coordinator assignments" >> db_checks.txt
$DB -c "SELECT a.id, a.user_id FROM coordinator_assignments a LEFT JOIN users u ON a.user_id = u.id WHERE u.id IS NULL;" >> db_checks.txt


-- Master seed file. Run with: make seed
-- Designed to run on a freshly migrated database.
-- All individual files are idempotent (ON CONFLICT DO NOTHING) except 06_jobs.sql,
-- which skips entirely if jobs already exist.

\i seeds/01_roles.sql
\i seeds/02_users.sql
\i seeds/03_companies.sql
\i seeds/04_departments.sql
\i seeds/05_skills.sql
\i seeds/06_jobs.sql
\i seeds/07_candidates.sql
\i seeds/08_candidate_skills.sql
\i seeds/09_job_skills.sql
\i seeds/10_applications.sql

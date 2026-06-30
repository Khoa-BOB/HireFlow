-- Seeds job_skills for every job inserted in 06_jobs.sql.
-- is_required = true means the skill is mandatory; false means nice-to-have.
DO $$
DECLARE
    go_engineer_id   UUID;
    frontend_id      UUID;
    devops_id        UUID;
    python_eng_id    UUID;
    platform_eng_id  UUID;
BEGIN
    SELECT id INTO go_engineer_id  FROM jobs WHERE title = 'Senior Go Engineer'     LIMIT 1;
    SELECT id INTO frontend_id     FROM jobs WHERE title = 'Frontend Engineer'      LIMIT 1;
    SELECT id INTO devops_id       FROM jobs WHERE title = 'DevOps Engineer'        LIMIT 1;
    SELECT id INTO python_eng_id   FROM jobs WHERE title = 'Backend Python Engineer' LIMIT 1;
    SELECT id INTO platform_eng_id FROM jobs WHERE title = 'Platform Engineer'      LIMIT 1;

    -- Senior Go Engineer
    INSERT INTO job_skills (job_id, skill_id, is_required) VALUES
        (go_engineer_id, (SELECT id FROM skills WHERE name = 'Go'),          true),
        (go_engineer_id, (SELECT id FROM skills WHERE name = 'PostgreSQL'),  true),
        (go_engineer_id, (SELECT id FROM skills WHERE name = 'REST API'),    true),
        (go_engineer_id, (SELECT id FROM skills WHERE name = 'Docker'),      true),
        (go_engineer_id, (SELECT id FROM skills WHERE name = 'Git'),         true),
        (go_engineer_id, (SELECT id FROM skills WHERE name = 'Redis'),       false),
        (go_engineer_id, (SELECT id FROM skills WHERE name = 'Kubernetes'),  false)
    ON CONFLICT (job_id, skill_id) DO NOTHING;

    -- Frontend Engineer
    INSERT INTO job_skills (job_id, skill_id, is_required) VALUES
        (frontend_id, (SELECT id FROM skills WHERE name = 'React'),        true),
        (frontend_id, (SELECT id FROM skills WHERE name = 'TypeScript'),   true),
        (frontend_id, (SELECT id FROM skills WHERE name = 'JavaScript'),   true),
        (frontend_id, (SELECT id FROM skills WHERE name = 'Git'),          true),
        (frontend_id, (SELECT id FROM skills WHERE name = 'Communication'), false)
    ON CONFLICT (job_id, skill_id) DO NOTHING;

    -- DevOps Engineer
    INSERT INTO job_skills (job_id, skill_id, is_required) VALUES
        (devops_id, (SELECT id FROM skills WHERE name = 'Docker'),     true),
        (devops_id, (SELECT id FROM skills WHERE name = 'Kubernetes'), true),
        (devops_id, (SELECT id FROM skills WHERE name = 'AWS'),        true),
        (devops_id, (SELECT id FROM skills WHERE name = 'Git'),        true),
        (devops_id, (SELECT id FROM skills WHERE name = 'Python'),     false)
    ON CONFLICT (job_id, skill_id) DO NOTHING;

    -- Backend Python Engineer
    INSERT INTO job_skills (job_id, skill_id, is_required) VALUES
        (python_eng_id, (SELECT id FROM skills WHERE name = 'Python'),     true),
        (python_eng_id, (SELECT id FROM skills WHERE name = 'SQL'),        true),
        (python_eng_id, (SELECT id FROM skills WHERE name = 'PostgreSQL'), true),
        (python_eng_id, (SELECT id FROM skills WHERE name = 'REST API'),   true),
        (python_eng_id, (SELECT id FROM skills WHERE name = 'Docker'),     true)
    ON CONFLICT (job_id, skill_id) DO NOTHING;

    -- Platform Engineer
    INSERT INTO job_skills (job_id, skill_id, is_required) VALUES
        (platform_eng_id, (SELECT id FROM skills WHERE name = 'Go'),         true),
        (platform_eng_id, (SELECT id FROM skills WHERE name = 'PostgreSQL'), true),
        (platform_eng_id, (SELECT id FROM skills WHERE name = 'Redis'),      true),
        (platform_eng_id, (SELECT id FROM skills WHERE name = 'Kubernetes'), true),
        (platform_eng_id, (SELECT id FROM skills WHERE name = 'AWS'),        true),
        (platform_eng_id, (SELECT id FROM skills WHERE name = 'gRPC'),       false)
    ON CONFLICT (job_id, skill_id) DO NOTHING;
END $$;

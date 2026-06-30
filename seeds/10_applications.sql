INSERT INTO applications (candidate_id, job_id, status, cover_letter, applied_at)
VALUES
    (
        (SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'bob@hireflow.dev'),
        (SELECT id FROM jobs WHERE title = 'Senior Go Engineer' LIMIT 1),
        'applied',
        'I have been writing Go professionally for four years and am excited about TechCorp''s focus on developer tooling. I would love to contribute to your core platform team.',
        now() - INTERVAL '5 days'
    ),
    (
        (SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'bob@hireflow.dev'),
        (SELECT id FROM jobs WHERE title = 'Platform Engineer' LIMIT 1),
        'saved',
        NULL,
        now() - INTERVAL '1 day'
    ),
    (
        (SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'diana@hireflow.dev'),
        (SELECT id FROM jobs WHERE title = 'Frontend Engineer' LIMIT 1),
        'screening',
        'React and TypeScript have been my primary stack for three years. I am particularly drawn to TechCorp''s commitment to design quality and would bring strong accessibility expertise to the team.',
        now() - INTERVAL '10 days'
    ),
    (
        (SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'eve@hireflow.dev'),
        (SELECT id FROM jobs WHERE title = 'Frontend Engineer' LIMIT 1),
        'applied',
        'I am a full-stack engineer looking to deepen my frontend focus. My Node.js background gives me strong empathy for API consumers, and I am eager to grow my React skills further.',
        now() - INTERVAL '3 days'
    ),
    (
        (SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'eve@hireflow.dev'),
        (SELECT id FROM jobs WHERE title = 'Backend Python Engineer' LIMIT 1),
        'applied',
        'I have built REST services in Node.js and am actively learning Python. I am drawn to HealthPlus''s mission and am excited to contribute to health data infrastructure.',
        now() - INTERVAL '2 days'
    )
ON CONFLICT (candidate_id, job_id) DO NOTHING;

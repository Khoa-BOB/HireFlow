-- proficiency_level: 1 (beginner) → 5 (expert)
INSERT INTO candidate_skills (candidate_id, skill_id, proficiency_level)
VALUES
    -- Bob: Go backend engineer
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'bob@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'Go'),              5),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'bob@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'Python'),          4),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'bob@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'SQL'),             4),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'bob@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'PostgreSQL'),      4),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'bob@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'REST API'),        5),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'bob@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'Docker'),          3),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'bob@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'Git'),             5),

    -- Diana: frontend engineer
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'diana@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'JavaScript'),      5),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'diana@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'TypeScript'),      5),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'diana@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'React'),           5),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'diana@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'Node.js'),         3),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'diana@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'Git'),             4),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'diana@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'Communication'),   5),

    -- Eve: full-stack engineer
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'eve@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'JavaScript'),      4),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'eve@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'TypeScript'),      3),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'eve@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'Node.js'),         5),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'eve@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'React'),           4),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'eve@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'PostgreSQL'),      4),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'eve@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'Docker'),          3),
    ((SELECT c.id FROM candidates c JOIN users u ON c.user_id = u.id WHERE u.email = 'eve@hireflow.dev'),
     (SELECT id FROM skills WHERE name = 'Git'),             4)
ON CONFLICT (candidate_id, skill_id) DO NOTHING;

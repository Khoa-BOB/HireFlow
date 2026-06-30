INSERT INTO skills (name, category)
VALUES
    ('Go',              'Programming Language'),
    ('Python',          'Programming Language'),
    ('JavaScript',      'Programming Language'),
    ('TypeScript',      'Programming Language'),
    ('Java',            'Programming Language'),
    ('SQL',             'Database'),
    ('PostgreSQL',      'Database'),
    ('Redis',           'Database'),
    ('React',           'Frontend'),
    ('Node.js',         'Backend'),
    ('REST API',        'Backend'),
    ('gRPC',            'Backend'),
    ('Docker',          'DevOps'),
    ('Kubernetes',      'DevOps'),
    ('AWS',             'Cloud'),
    ('Git',             'Tool'),
    ('Communication',   'Soft Skill'),
    ('Problem Solving', 'Soft Skill'),
    ('Leadership',      'Soft Skill')
ON CONFLICT (name) DO NOTHING;

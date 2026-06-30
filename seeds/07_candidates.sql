INSERT INTO candidates (user_id, bio, city, country, linkedin_url, portfolio_url)
VALUES
    (
        (SELECT id FROM users WHERE email = 'bob@hireflow.dev'),
        'Backend engineer with 4 years of experience building distributed systems in Go and Python. Passionate about clean APIs and database performance.',
        'Austin',
        'US',
        'https://linkedin.com/in/bobcandidate',
        'https://github.com/bobcandidate'
    ),
    (
        (SELECT id FROM users WHERE email = 'diana@hireflow.dev'),
        'Frontend engineer specializing in React and TypeScript. I care deeply about accessibility, design systems, and developer experience.',
        'Seattle',
        'US',
        'https://linkedin.com/in/dianalee',
        'https://dianalee.dev'
    ),
    (
        (SELECT id FROM users WHERE email = 'eve@hireflow.dev'),
        'Full-stack engineer with 3 years of experience across Node.js backends and React frontends. Recently exploring cloud infrastructure and Kubernetes.',
        'Chicago',
        'US',
        'https://linkedin.com/in/evechen',
        NULL
    )
ON CONFLICT (user_id) DO NOTHING;

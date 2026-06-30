INSERT INTO companies (name, description, industry, url, city, country)
VALUES
    (
        'TechCorp Inc',
        'A leading software company building developer tools and cloud infrastructure.',
        'Technology',
        'https://techcorp.example.com',
        'San Francisco',
        'US'
    ),
    (
        'HealthPlus',
        'Digital health platform improving patient outcomes through data-driven care.',
        'Healthcare',
        'https://healthplus.example.com',
        'Boston',
        'US'
    ),
    (
        'FinanceWorks',
        'Fintech company powering payment infrastructure for modern businesses.',
        'Finance',
        'https://financeworks.example.com',
        'New York',
        'US'
    )
ON CONFLICT (name) DO NOTHING;

-- Link staff users to TechCorp
UPDATE users
SET company_id = (SELECT id FROM companies WHERE name = 'TechCorp Inc')
WHERE email IN ('recruiter@hireflow.dev', 'manager@hireflow.dev')
  AND company_id IS NULL;

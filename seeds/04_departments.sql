INSERT INTO departments (company_id, name, description)
VALUES
    (
        (SELECT id FROM companies WHERE name = 'TechCorp Inc'),
        'Engineering',
        'Builds and maintains core platform products'
    ),
    (
        (SELECT id FROM companies WHERE name = 'TechCorp Inc'),
        'Product',
        'Defines product vision, strategy, and roadmap'
    ),
    (
        (SELECT id FROM companies WHERE name = 'HealthPlus'),
        'Engineering',
        'Develops healthcare software and data pipelines'
    ),
    (
        (SELECT id FROM companies WHERE name = 'HealthPlus'),
        'Operations',
        'Manages infrastructure, reliability, and cloud operations'
    ),
    (
        (SELECT id FROM companies WHERE name = 'FinanceWorks'),
        'Engineering',
        'Payment systems, platform APIs, and core services'
    ),
    (
        (SELECT id FROM companies WHERE name = 'FinanceWorks'),
        'Risk & Compliance',
        'Manages regulatory compliance and financial risk'
    )
ON CONFLICT (company_id, name) DO NOTHING;

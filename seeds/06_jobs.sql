-- Jobs have no natural unique key; this block skips insert if any jobs already exist.
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM jobs) THEN
        INSERT INTO jobs (department_id, posted_by, title, description, employment_type, status, salary_min, salary_max, currency)
        VALUES
            (
                (SELECT d.id FROM departments d JOIN companies c ON d.company_id = c.id
                 WHERE c.name = 'TechCorp Inc' AND d.name = 'Engineering'),
                (SELECT id FROM users WHERE email = 'recruiter@hireflow.dev'),
                'Senior Go Engineer',
                'We are looking for an experienced Go engineer to design and build core platform APIs. You will own backend services end-to-end and collaborate closely with product and infrastructure teams.',
                'Full-time',
                'open',
                120000, 160000, 'USD'
            ),
            (
                (SELECT d.id FROM departments d JOIN companies c ON d.company_id = c.id
                 WHERE c.name = 'TechCorp Inc' AND d.name = 'Engineering'),
                (SELECT id FROM users WHERE email = 'recruiter@hireflow.dev'),
                'Frontend Engineer',
                'Join our product team to build beautiful, performant interfaces using React and TypeScript. You will work closely with designers and backend engineers to deliver great user experiences.',
                'Full-time',
                'open',
                100000, 140000, 'USD'
            ),
            (
                (SELECT d.id FROM departments d JOIN companies c ON d.company_id = c.id
                 WHERE c.name = 'TechCorp Inc' AND d.name = 'Engineering'),
                (SELECT id FROM users WHERE email = 'recruiter@hireflow.dev'),
                'DevOps Engineer',
                'Own our CI/CD pipelines, Kubernetes clusters, and cloud infrastructure on AWS. You will drive reliability, scalability, and developer productivity across engineering.',
                'Full-time',
                'open',
                110000, 150000, 'USD'
            ),
            (
                (SELECT d.id FROM departments d JOIN companies c ON d.company_id = c.id
                 WHERE c.name = 'HealthPlus' AND d.name = 'Engineering'),
                NULL,
                'Backend Python Engineer',
                'Build scalable microservices for our health data processing platform. Experience with healthcare data standards (HL7, FHIR) is a plus.',
                'Full-time',
                'open',
                90000, 130000, 'USD'
            ),
            (
                (SELECT d.id FROM departments d JOIN companies c ON d.company_id = c.id
                 WHERE c.name = 'FinanceWorks' AND d.name = 'Engineering'),
                NULL,
                'Platform Engineer',
                'Design and scale payment infrastructure handling millions of daily transactions. You will work on low-latency, high-availability systems critical to our business.',
                'Full-time',
                'open',
                130000, 180000, 'USD'
            );
    END IF;
END $$;

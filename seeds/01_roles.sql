INSERT INTO roles (name, description)
VALUES
    ('admin',          'System administrator with full access'),
    ('recruiter',      'Manages job postings and candidates'),
    ('hiring_manager', 'Reviews shortlisted candidates and makes hiring decisions'),
    ('candidate',      'Job seeker applying for positions')
ON CONFLICT (name) DO NOTHING;

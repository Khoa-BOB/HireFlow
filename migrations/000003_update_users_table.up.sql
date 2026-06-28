INSERT INTO roles (name, description) VALUES
    ('candidate',       'Job seeker applying for positions'),
    ('recruiter',       'Manages job postings and candidates'),
    ('hiring_manager',  'Reviews shortlisted candidates and makes hiring decisions')
ON CONFLICT (name) DO NOTHING;

ALTER TABLE users
ADD COLUMN role_id UUID REFERENCES roles(id);

UPDATE users
SET role_id = (SELECT id FROM roles WHERE name = 'candidate')
WHERE role_id IS NULL;

ALTER TABLE users
ALTER COLUMN role_id SET NOT NULL;

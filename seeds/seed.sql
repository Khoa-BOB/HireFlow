-- Seed users (passwords shown in comments)
INSERT INTO users (email, password_hash, full_name, role_id)
VALUES
    (
        'admin@hireflow.dev',
        '$2a$10$pITSgjKHNWRQxjfCCKrWV.nKNHEvCcMWLaY7QKtvlCDBo0Mx4dDly', -- admin123
        'Admin User',
        (SELECT id FROM roles WHERE name = 'admin')
    ),
    (
        'recruiter@hireflow.dev',
        '$2a$10$Gw4VyIzYE6JC1QgY4TXnAu7WoJ49AIEBoT3HKqJ7ESJbYh9qtVlaa', -- recruiter123
        'Alice Recruiter',
        (SELECT id FROM roles WHERE name = 'recruiter')
    ),
    (
        'candidate@hireflow.dev',
        '$2a$10$UHqvaYd7wW9VHkyB765F1ulJJx7AtJgc9kRWKOC3OWbX1o.CvcHee', -- candidate123
        'Bob Candidate',
        (SELECT id FROM roles WHERE name = 'candidate')
    )
ON CONFLICT (email) DO NOTHING;

-- Passwords: admin123 / recruiter123 / candidate123
INSERT INTO users (email, password_hash, full_name, role_id)
VALUES
    (
        'admin@hireflow.dev',
        '$2a$10$pITSgjKHNWRQxjfCCKrWV.nKNHEvCcMWLaY7QKtvlCDBo0Mx4dDly',
        'Admin User',
        (SELECT id FROM roles WHERE name = 'admin')
    ),
    (
        'recruiter@hireflow.dev',
        '$2a$10$Gw4VyIzYE6JC1QgY4TXnAu7WoJ49AIEBoT3HKqJ7ESJbYh9qtVlaa',
        'Alice Recruiter',
        (SELECT id FROM roles WHERE name = 'recruiter')
    ),
    (
        'manager@hireflow.dev',
        '$2a$10$Gw4VyIzYE6JC1QgY4TXnAu7WoJ49AIEBoT3HKqJ7ESJbYh9qtVlaa',
        'Charlie Manager',
        (SELECT id FROM roles WHERE name = 'hiring_manager')
    ),
    (
        'bob@hireflow.dev',
        '$2a$10$UHqvaYd7wW9VHkyB765F1ulJJx7AtJgc9kRWKOC3OWbX1o.CvcHee',
        'Bob Candidate',
        (SELECT id FROM roles WHERE name = 'candidate')
    ),
    (
        'diana@hireflow.dev',
        '$2a$10$UHqvaYd7wW9VHkyB765F1ulJJx7AtJgc9kRWKOC3OWbX1o.CvcHee',
        'Diana Lee',
        (SELECT id FROM roles WHERE name = 'candidate')
    ),
    (
        'eve@hireflow.dev',
        '$2a$10$UHqvaYd7wW9VHkyB765F1ulJJx7AtJgc9kRWKOC3OWbX1o.CvcHee',
        'Eve Chen',
        (SELECT id FROM roles WHERE name = 'candidate')
    )
ON CONFLICT (email) DO NOTHING;

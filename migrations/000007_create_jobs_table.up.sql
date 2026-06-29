CREATE TABLE jobs(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    department_id UUID NOT NULL REFERENCES departments(id) ON DELETE CASCADE,
    posted_by UUID REFERENCES users(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    employment_type VARCHAR(50),
    status VARCHAR(20) NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'open', 'closed', 'archived')),
    salary_min NUMERIC(12,2) CHECK (salary_min >= 0),
    salary_max NUMERIC(12,2) CHECK (salary_max >= 0),
    currency CHAR(3) NOT NULL DEFAULT 'USD',
    CHECK (salary_max >= salary_min),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
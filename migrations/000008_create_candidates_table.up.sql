CREATE TABLE candidates(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    bio TEXT,
    city VARCHAR(100),
    country VARCHAR(100),
    linkedin_url VARCHAR(2048),
    portfolio_url VARCHAR(2048),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
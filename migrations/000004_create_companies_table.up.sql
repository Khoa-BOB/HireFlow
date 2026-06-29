CREATE TABLE companies(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR (255) NOT NULL UNIQUE,
    description TEXT,
    industry VARCHAR(100),
    url VARCHAR(2048),
    address TEXT,
    city VARCHAR(100),
    country VARCHAR(100),
    zipcode VARCHAR(20),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);


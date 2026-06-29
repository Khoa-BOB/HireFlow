CREATE TABLE applications(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    candidate_id UUID NOT NULL REFERENCES candidates(id) ON DELETE CASCADE,
    job_id UUID NOT NULL REFERENCES jobs(id) ON DELETE CASCADE,

    status VARCHAR(50) NOT NULL DEFAULT 'saved' CHECK (status IN ('saved', 'applied', 'screening', 'interview', 'offer', 'rejected', 'accepted', 'withdrawn')),
    cover_letter TEXT,
    applied_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    UNIQUE(candidate_id, job_id)
);
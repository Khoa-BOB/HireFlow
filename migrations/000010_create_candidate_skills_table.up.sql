CREATE TABLE candidate_skills(
    candidate_id UUID NOT NULL REFERENCES candidates(id) ON DELETE CASCADE,
    skill_id UUID NOT NULL REFERENCES skills(id) ON DELETE CASCADE,
    proficiency_level SMALLINT CHECK (proficiency_level BETWEEN 1 AND 5),
    PRIMARY KEY (candidate_id, skill_id)
);

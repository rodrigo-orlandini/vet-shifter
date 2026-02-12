CREATE TABLE ratings (
  id UUID PRIMARY KEY,
  shift_id UUID NOT NULL REFERENCES shifts(id) ON DELETE CASCADE,
  from_company_id UUID REFERENCES companies(id) ON DELETE SET NULL,
  from_vet_id UUID REFERENCES shift_vets(id) ON DELETE SET NULL,
  to_company_id UUID REFERENCES companies(id) ON DELETE SET NULL,
  to_vet_id UUID REFERENCES shift_vets(id) ON DELETE SET NULL,
  score INT NOT NULL CHECK (score >= 1 AND score <= 5),
  comment TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_ratings_shift_id ON ratings(shift_id);
CREATE INDEX idx_ratings_to_company_id ON ratings(to_company_id);
CREATE INDEX idx_ratings_to_vet_id ON ratings(to_vet_id);

CREATE TABLE shift_candidacies (
  id UUID PRIMARY KEY,
  shift_id UUID NOT NULL REFERENCES shifts(id) ON DELETE CASCADE,
  shift_vet_id UUID NOT NULL REFERENCES shift_vets(id) ON DELETE CASCADE,
  status VARCHAR(20) NOT NULL DEFAULT 'pending',
  invited_by_clinic BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(shift_id, shift_vet_id)
);

CREATE INDEX idx_shift_candidacies_shift_id ON shift_candidacies(shift_id);
CREATE INDEX idx_shift_candidacies_shift_vet_id ON shift_candidacies(shift_vet_id);

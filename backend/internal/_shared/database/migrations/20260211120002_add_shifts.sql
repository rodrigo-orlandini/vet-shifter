CREATE TABLE shifts (
  id UUID PRIMARY KEY,
  company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  starts_at TIMESTAMP NOT NULL,
  ends_at TIMESTAMP NOT NULL,
  type VARCHAR(50) NOT NULL,
  offered_value_cents BIGINT NOT NULL,
  requirements TEXT,
  description TEXT,
  location VARCHAR(500),
  status VARCHAR(20) NOT NULL DEFAULT 'open',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_shifts_company_id ON shifts(company_id);
CREATE INDEX idx_shifts_starts_at ON shifts(starts_at);
CREATE INDEX idx_shifts_status ON shifts(status);
CREATE INDEX idx_shifts_type ON shifts(type);

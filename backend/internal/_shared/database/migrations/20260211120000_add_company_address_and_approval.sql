-- Add address and approval to companies; LGPD consent to company_owners
ALTER TABLE companies
  ADD COLUMN street VARCHAR(255),
  ADD COLUMN number VARCHAR(20),
  ADD COLUMN city VARCHAR(100),
  ADD COLUMN state VARCHAR(2),
  ADD COLUMN zip_code VARCHAR(10),
  ADD COLUMN approval_status VARCHAR(20) NOT NULL DEFAULT 'pending';

ALTER TABLE company_owners
  ADD COLUMN consent_lgpd_at TIMESTAMP;

CREATE INDEX idx_companies_approval_status ON companies(approval_status);

CREATE TABLE shift_vets (
  id UUID PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  phone VARCHAR(20) NOT NULL,
  password VARCHAR(255) NOT NULL,
  full_name VARCHAR(255) NOT NULL,
  cpf VARCHAR(11) NOT NULL UNIQUE,
  crmv_number VARCHAR(20) NOT NULL,
  crmv_state VARCHAR(2) NOT NULL,
  specialties TEXT[] DEFAULT '{}',
  approval_status VARCHAR(20) NOT NULL DEFAULT 'pending',
  consent_lgpd_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_shift_vets_email ON shift_vets(email);
CREATE INDEX idx_shift_vets_cpf ON shift_vets(cpf);
CREATE INDEX idx_shift_vets_approval_status ON shift_vets(approval_status);

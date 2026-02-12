-- Full schema for sqlc codegen (matches migrations state)
CREATE TABLE companies (
  id UUID PRIMARY KEY,
  cnpj VARCHAR(14) NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL,
  street VARCHAR(255),
  number VARCHAR(20),
  city VARCHAR(100),
  state VARCHAR(2),
  zip_code VARCHAR(10),
  approval_status VARCHAR(20) NOT NULL DEFAULT 'pending',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE company_owners (
  id UUID PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  phone VARCHAR(20) NOT NULL,
  password VARCHAR(255) NOT NULL,
  company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  consent_lgpd_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

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

CREATE TABLE shift_candidacies (
  id UUID PRIMARY KEY,
  shift_id UUID NOT NULL REFERENCES shifts(id) ON DELETE CASCADE,
  shift_vet_id UUID NOT NULL REFERENCES shift_vets(id) ON DELETE CASCADE,
  status VARCHAR(20) NOT NULL DEFAULT 'pending',
  invited_by_clinic BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(shift_id, shift_vet_id)
);

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

CREATE TYPE account_status AS ENUM (
    'pending_document_approval',
    'complete'
);

CREATE TYPE veterinary_specialty AS ENUM (
    'general_practice',
    'felines',
    'wildlife',
    'dermatology',
    'cardiology',
    'nephrology',
    'urology',
    'endocrinology',
    'gastroenterology',
    'neurology',
    'orthopedics',
    'dentistry',
    'ophthalmology',
    'ultrasound',
    'pathology',
    'anesthesiology',
    'icu',
    'oncology',
    'physiotherapy',
    'behavioral'
);

CREATE TABLE IF NOT EXISTS companies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cnpj VARCHAR(14) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    street VARCHAR(255),
    number VARCHAR(20),
    city VARCHAR(100),
    state VARCHAR(2),
    zip_code VARCHAR(10),
    approval_status account_status NOT NULL DEFAULT 'pending_document_approval',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS company_owners (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    consent_lgpd_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS shift_veterinaries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    cpf VARCHAR(11) NOT NULL UNIQUE,
    crmv_number VARCHAR(20) NOT NULL,
    crmv_state VARCHAR(2) NOT NULL,
    specialties veterinary_specialty[] NOT NULL DEFAULT '{}',
    approval_status account_status NOT NULL DEFAULT 'pending_document_approval',
    consent_lgpd_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TYPE user_type AS ENUM (
    'company_owner',
    'shift_veterinary'
);

CREATE TABLE IF NOT EXISTS password_reset_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    token VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL,
    user_type user_type NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    used_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_password_reset_tokens_token ON password_reset_tokens(token);
CREATE INDEX IF NOT EXISTS idx_password_reset_tokens_expires_at ON password_reset_tokens(expires_at);

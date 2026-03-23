CREATE TYPE account_status AS ENUM (
    'pending',
    'approved',
    'rejected',
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

CREATE TABLE companies (
    id UUID PRIMARY KEY,
    cnpj VARCHAR(14) NOT NULL,
    name VARCHAR(255) NOT NULL,
    approval_status account_status NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE addresses (
    id UUID PRIMARY KEY,
    company_id UUID NOT NULL,
    street VARCHAR(255),
    number VARCHAR(50),
    city VARCHAR(100),
    state VARCHAR(2),
    zip_code VARCHAR(10),
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE company_owners (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(30) NOT NULL,
    password VARCHAR(255) NOT NULL,
    company_id UUID NOT NULL,
    consent_lgpd_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE shift_veterinaries (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(30) NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    cpf VARCHAR(11) NOT NULL,
    crmv_number VARCHAR(30) NOT NULL,
    crmv_state VARCHAR(2) NOT NULL,
    specialties veterinary_specialty[] NOT NULL,
    approval_status account_status NOT NULL,
    consent_lgpd_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TYPE user_type AS ENUM (
    'company_owner',
    'shift_veterinary'
);

CREATE TABLE password_reset_tokens (
    id UUID PRIMARY KEY,
    token VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    user_type user_type NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    used_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL
);

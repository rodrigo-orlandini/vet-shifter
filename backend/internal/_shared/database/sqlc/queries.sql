-- Companies
-- name: CreateCompany :one
INSERT INTO companies (id, cnpj, name, street, number, city, state, zip_code, approval_status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, cnpj, name, street, number, city, state, zip_code, approval_status, created_at;

-- name: FindCompanyByCnpj :one
SELECT id, cnpj, name, street, number, city, state, zip_code, approval_status, created_at FROM companies WHERE cnpj = $1;

-- name: FindCompanyByID :one
SELECT id, cnpj, name, street, number, city, state, zip_code, approval_status, created_at FROM companies WHERE id = $1;

-- Company owners
-- name: FindCompanyOwnerByEmail :one
SELECT id, email, phone, password, company_id, consent_lgpd_at, created_at FROM company_owners WHERE email = $1;

-- name: FindCompanyOwnerByID :one
SELECT id, email, phone, password, company_id, consent_lgpd_at, created_at FROM company_owners WHERE id = $1;

-- name: RegisterCompanyOwner :exec
INSERT INTO company_owners (id, email, phone, password, company_id, consent_lgpd_at)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateCompanyOwnerPassword :exec
UPDATE company_owners SET password = $2 WHERE id = $1;

-- Shift veterinaries
-- name: CreateShiftVeterinary :one
INSERT INTO shift_veterinaries (id, email, phone, password, full_name, cpf, crmv_number, crmv_state, specialties, approval_status, consent_lgpd_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id, email, phone, password, full_name, cpf, crmv_number, crmv_state, specialties, approval_status, consent_lgpd_at, created_at;

-- name: FindShiftVeterinaryByCpf :one
SELECT id, email, phone, password, full_name, cpf, crmv_number, crmv_state, specialties, approval_status, consent_lgpd_at, created_at FROM shift_veterinaries WHERE cpf = $1;

-- name: FindShiftVeterinaryByEmail :one
SELECT id, email, phone, password, full_name, cpf, crmv_number, crmv_state, specialties, approval_status, consent_lgpd_at, created_at FROM shift_veterinaries WHERE email = $1;

-- name: FindShiftVeterinaryByID :one
SELECT id, email, phone, password, full_name, cpf, crmv_number, crmv_state, specialties, approval_status, consent_lgpd_at, created_at FROM shift_veterinaries WHERE id = $1;

-- name: UpdateShiftVeterinaryPassword :exec
UPDATE shift_veterinaries SET password = $2 WHERE id = $1;

-- Password reset tokens
-- name: CreatePasswordResetToken :one
INSERT INTO password_reset_tokens (id, token, email, user_type, expires_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, token, email, user_type, expires_at, used_at, created_at;

-- name: GetPasswordResetTokenByToken :one
SELECT id, token, email, user_type, expires_at, used_at, created_at FROM password_reset_tokens WHERE token = $1;

-- name: MarkPasswordResetTokenUsed :exec
UPDATE password_reset_tokens SET used_at = NOW() WHERE id = $1;

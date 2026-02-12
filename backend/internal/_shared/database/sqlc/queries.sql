-- Companies (with address and approval)
-- name: CreateCompany :one
INSERT INTO companies (id, cnpj, name, street, number, city, state, zip_code, approval_status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: FindCompanyByCnpj :one
SELECT * FROM companies WHERE cnpj = $1;

-- name: FindCompanyByID :one
SELECT * FROM companies WHERE id = $1;

-- Company owners
-- name: RegisterCompanyOwner :exec
INSERT INTO company_owners (id, email, phone, password, company_id, consent_lgpd_at)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: FindCompanyOwnerByEmail :one
SELECT * FROM company_owners WHERE email = $1;

-- name: FindCompanyOwnerByID :one
SELECT * FROM company_owners WHERE id = $1;

-- Shift vets (plantonistas)
-- name: CreateShiftVet :one
INSERT INTO shift_vets (id, email, phone, password, full_name, cpf, crmv_number, crmv_state, specialties, approval_status, consent_lgpd_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: FindShiftVetByEmail :one
SELECT * FROM shift_vets WHERE email = $1;

-- name: FindShiftVetByCpf :one
SELECT * FROM shift_vets WHERE cpf = $1;

-- name: FindShiftVetByID :one
SELECT * FROM shift_vets WHERE id = $1;

-- Shifts
-- name: CreateShift :one
INSERT INTO shifts (id, company_id, starts_at, ends_at, type, offered_value_cents, requirements, description, location, status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetShiftByID :one
SELECT * FROM shifts WHERE id = $1;

-- name: ListShifts :many
SELECT * FROM shifts
WHERE
  ($1 = '' OR status = $1)
  AND ($2 = '' OR type = $2)
  AND ($3 = '00000000-0000-0000-0000-000000000000' OR company_id = $3)
  AND (starts_at >= $4)
  AND (starts_at <= $5)
ORDER BY starts_at ASC
LIMIT $6 OFFSET $7;

-- Shift candidacies
-- name: CreateShiftCandidacy :one
INSERT INTO shift_candidacies (id, shift_id, shift_vet_id, status, invited_by_clinic)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: FindCandidacyByShiftAndVet :one
SELECT * FROM shift_candidacies WHERE shift_id = $1 AND shift_vet_id = $2;

-- name: ListCandidaciesByShift :many
SELECT * FROM shift_candidacies WHERE shift_id = $1 ORDER BY created_at DESC;

-- name: UpdateCandidacyStatus :exec
UPDATE shift_candidacies SET status = $2 WHERE id = $1;

-- name: GetCandidacyByID :one
SELECT * FROM shift_candidacies WHERE id = $1;

-- When accepting a candidacy we need to update shift status
-- name: UpdateShiftStatus :exec
UPDATE shifts SET status = $2 WHERE id = $1;

-- Ratings
-- name: CreateRating :one
INSERT INTO ratings (id, shift_id, from_company_id, from_vet_id, to_company_id, to_vet_id, score, comment)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: ListRatingsByCompany :many
SELECT * FROM ratings WHERE to_company_id = $1 ORDER BY created_at DESC;

-- name: ListRatingsByVet :many
SELECT * FROM ratings WHERE to_vet_id = $1 ORDER BY created_at DESC;

-- name: ExistsRatingFromCompanyForShift :one
SELECT EXISTS(SELECT 1 FROM ratings WHERE shift_id = $1 AND from_company_id = $2) AS exists;

-- name: ExistsRatingFromVetForShift :one
SELECT EXISTS(SELECT 1 FROM ratings WHERE shift_id = $1 AND from_vet_id = $2) AS exists;

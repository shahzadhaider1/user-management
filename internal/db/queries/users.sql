-- name: CreateUser :one
INSERT INTO users (name, phone_number)
VALUES ($1, $2)
RETURNING id;

-- name: GetUserByPhoneNumber :one
SELECT * FROM users WHERE phone_number = $1;

-- name: UpdateOTP :exec
UPDATE users
SET otp = $1, otp_expiration_time = $2
WHERE phone_number = $3;

-- name: VerifyOTP :one
SELECT * FROM users WHERE phone_number = $1 AND otp = $2 AND otp_expiration_time > NOW();

-- queries.sql

-- name: CreateFact :one
INSERT INTO facts (
    bedroom, bathroom, plot_area, built_up_area, view, furnished, ownership, sc_currency_id, unit_of_measure
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING id;

-- name: GetFact :one
SELECT
    id, bedroom, bathroom, plot_area, built_up_area, view, furnished, ownership, sc_currency_id, unit_of_measure
FROM facts
WHERE id = $1;

-- name: ListFacts :many
SELECT
    id, bedroom, bathroom, plot_area, built_up_area, view, furnished, ownership, sc_currency_id, unit_of_measure
FROM facts;

-- name: UpdateFact :one
-- queries.sql

-- name: UpdateFact :one
UPDATE facts
SET
    bedroom = COALESCE(NULLIF($1, '{}'), bedroom),
    bathroom = COALESCE(NULLIF($2, '{}'), bathroom),
    plot_area = COALESCE($3, plot_area), -- Using COALESCE to handle null values
    built_up_area = COALESCE($4, built_up_area), -- Using COALESCE to handle null values
    view = COALESCE(NULLIF($5, '{}'), view),
    furnished = COALESCE($6, furnished),
    ownership = COALESCE($7, ownership),
    sc_currency_id = COALESCE($8, sc_currency_id),
    unit_of_measure = COALESCE($9, unit_of_measure)
WHERE id = $10
RETURNING id;

-- name: DeleteFact :exec
DELETE FROM facts
WHERE id = $1;

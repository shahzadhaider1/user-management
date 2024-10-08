// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createFact = `-- name: CreateFact :one

INSERT INTO facts (
    bedroom, bathroom, plot_area, built_up_area, view, furnished, ownership, sc_currency_id, unit_of_measure
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING id
`

type CreateFactParams struct {
	Bedroom       []string
	Bathroom      []int64
	PlotArea      sql.NullFloat64
	BuiltUpArea   sql.NullFloat64
	View          []int64
	Furnished     sql.NullInt64
	Ownership     sql.NullInt64
	ScCurrencyID  sql.NullString
	UnitOfMeasure sql.NullString
}

// queries.sql
func (q *Queries) CreateFact(ctx context.Context, arg CreateFactParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createFact,
		pq.Array(arg.Bedroom),
		pq.Array(arg.Bathroom),
		arg.PlotArea,
		arg.BuiltUpArea,
		pq.Array(arg.View),
		arg.Furnished,
		arg.Ownership,
		arg.ScCurrencyID,
		arg.UnitOfMeasure,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, phone_number)
VALUES ($1, $2)
RETURNING id
`

type CreateUserParams struct {
	Name        string
	PhoneNumber string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.PhoneNumber)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteFact = `-- name: DeleteFact :exec
DELETE FROM facts
WHERE id = $1
`

func (q *Queries) DeleteFact(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteFact, id)
	return err
}

const getFact = `-- name: GetFact :one
SELECT
    id, bedroom, bathroom, plot_area, built_up_area, view, furnished, ownership, sc_currency_id, unit_of_measure
FROM facts
WHERE id = $1
`

func (q *Queries) GetFact(ctx context.Context, id int64) (Fact, error) {
	row := q.db.QueryRowContext(ctx, getFact, id)
	var i Fact
	err := row.Scan(
		&i.ID,
		pq.Array(&i.Bedroom),
		pq.Array(&i.Bathroom),
		&i.PlotArea,
		&i.BuiltUpArea,
		pq.Array(&i.View),
		&i.Furnished,
		&i.Ownership,
		&i.ScCurrencyID,
		&i.UnitOfMeasure,
	)
	return i, err
}

const getUserByPhoneNumber = `-- name: GetUserByPhoneNumber :one
SELECT id, name, phone_number, otp, otp_expiration_time FROM users WHERE phone_number = $1
`

func (q *Queries) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByPhoneNumber, phoneNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PhoneNumber,
		&i.Otp,
		&i.OtpExpirationTime,
	)
	return i, err
}

const listFacts = `-- name: ListFacts :many
SELECT
    id, bedroom, bathroom, plot_area, built_up_area, view, furnished, ownership, sc_currency_id, unit_of_measure
FROM facts
`

func (q *Queries) ListFacts(ctx context.Context) ([]Fact, error) {
	rows, err := q.db.QueryContext(ctx, listFacts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Fact
	for rows.Next() {
		var i Fact
		if err := rows.Scan(
			&i.ID,
			pq.Array(&i.Bedroom),
			pq.Array(&i.Bathroom),
			&i.PlotArea,
			&i.BuiltUpArea,
			pq.Array(&i.View),
			&i.Furnished,
			&i.Ownership,
			&i.ScCurrencyID,
			&i.UnitOfMeasure,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFact = `-- name: UpdateFact :one

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
RETURNING id
`

type UpdateFactParams struct {
	Column1       interface{}
	Column2       interface{}
	PlotArea      sql.NullFloat64
	BuiltUpArea   sql.NullFloat64
	Column5       interface{}
	Furnished     sql.NullInt64
	Ownership     sql.NullInt64
	ScCurrencyID  sql.NullString
	UnitOfMeasure sql.NullString
	ID            int64
}

// queries.sql
func (q *Queries) UpdateFact(ctx context.Context, arg UpdateFactParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateFact,
		arg.Column1,
		arg.Column2,
		arg.PlotArea,
		arg.BuiltUpArea,
		arg.Column5,
		arg.Furnished,
		arg.Ownership,
		arg.ScCurrencyID,
		arg.UnitOfMeasure,
		arg.ID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateOTP = `-- name: UpdateOTP :exec
UPDATE users
SET otp = $1, otp_expiration_time = $2
WHERE phone_number = $3
`

type UpdateOTPParams struct {
	Otp               sql.NullString
	OtpExpirationTime sql.NullTime
	PhoneNumber       string
}

func (q *Queries) UpdateOTP(ctx context.Context, arg UpdateOTPParams) error {
	_, err := q.db.ExecContext(ctx, updateOTP, arg.Otp, arg.OtpExpirationTime, arg.PhoneNumber)
	return err
}

const verifyOTP = `-- name: VerifyOTP :one
SELECT id, name, phone_number, otp, otp_expiration_time FROM users WHERE phone_number = $1 AND otp = $2 AND otp_expiration_time > NOW()
`

type VerifyOTPParams struct {
	PhoneNumber string
	Otp         sql.NullString
}

func (q *Queries) VerifyOTP(ctx context.Context, arg VerifyOTPParams) (User, error) {
	row := q.db.QueryRowContext(ctx, verifyOTP, arg.PhoneNumber, arg.Otp)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PhoneNumber,
		&i.Otp,
		&i.OtpExpirationTime,
	)
	return i, err
}

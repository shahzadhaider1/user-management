package db

import (
	"context"
	"database/sql"
	"time"

	"user-management/internal/db"
)

// CreateUser creates a new user in the database
func (store *Store) CreateUser(ctx context.Context, name, phoneNumber string) (int32, error) {
	return store.Queries.CreateUser(ctx, db.CreateUserParams{
		Name:        name,
		PhoneNumber: phoneNumber,
	})
}

// GetUserByPhoneNumber retrieves a user by phone number
func (store *Store) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (db.User, error) {
	return store.Queries.GetUserByPhoneNumber(ctx, phoneNumber)
}

// UpdateOTP updates the OTP and its expiration time for a user
func (store *Store) UpdateOTP(ctx context.Context, phoneNumber, otp string, expiration time.Time) error {
	return store.Queries.UpdateOTP(ctx, db.UpdateOTPParams{
		Otp:               sql.NullString{String: otp, Valid: true},
		OtpExpirationTime: sql.NullTime{Time: expiration, Valid: true},
		PhoneNumber:       phoneNumber,
	})
}

// VerifyOTP verifies the OTP for a given phone number
func (store *Store) VerifyOTP(ctx context.Context, phoneNumber, otp string) (db.User, error) {
	return store.Queries.VerifyOTP(ctx, db.VerifyOTPParams{
		PhoneNumber: phoneNumber,
		Otp:         sql.NullString{String: otp, Valid: true},
	})
}

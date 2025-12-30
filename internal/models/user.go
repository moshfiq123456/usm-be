package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents the users table
type User struct {
	ID            uuid.UUID  `json:"id" db:"id"`
	Username      string     `json:"username" db:"username"`
	Email         string     `json:"email" db:"email"`
	PasswordHash  string     `json:"-" db:"password_hash"` // "-" means don't include in JSON
	EmailVerified bool       `json:"email_verified" db:"email_verified"`
	IsActive      bool       `json:"is_active" db:"is_active"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
	LastLoginAt   *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// UserProfile represents the user_profiles table
type UserProfile struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	UserID      uuid.UUID  `json:"user_id" db:"user_id"`
	FirstName   *string    `json:"first_name,omitempty" db:"first_name"`
	LastName    *string    `json:"last_name,omitempty" db:"last_name"`
	Phone       *string    `json:"phone,omitempty" db:"phone"`
	AvatarURL   *string    `json:"avatar_url,omitempty" db:"avatar_url"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty" db:"date_of_birth"`
	Bio         *string    `json:"bio,omitempty" db:"bio"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

// UserWithProfile combines User and UserProfile for API responses
type UserWithProfile struct {
	User
	Profile *UserProfile `json:"profile,omitempty"`
}
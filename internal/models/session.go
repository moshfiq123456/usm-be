package models

import (
	"time"

	"github.com/google/uuid"
)

// Session represents the sessions table
type Session struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	UserID         uuid.UUID  `json:"user_id" db:"user_id"`
	Token          string     `json:"token" db:"token"`
	RefreshToken   *string    `json:"refresh_token,omitempty" db:"refresh_token"`
	IPAddress      *string    `json:"ip_address,omitempty" db:"ip_address"`
	UserAgent      *string    `json:"user_agent,omitempty" db:"user_agent"`
	ExpiresAt      time.Time  `json:"expires_at" db:"expires_at"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	LastActivityAt time.Time  `json:"last_activity_at" db:"last_activity_at"`
}
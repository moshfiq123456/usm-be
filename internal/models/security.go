package models

import (
	"time"

	"github.com/google/uuid"
)

// LoginAttempt represents the login_attempts table
type LoginAttempt struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Email       string    `json:"email" db:"email"`
	IPAddress   string    `json:"ip_address" db:"ip_address"`
	Successful  bool      `json:"successful" db:"successful"`
	AttemptedAt time.Time `json:"attempted_at" db:"attempted_at"`
}

// AuditLog represents the audit_logs table
type AuditLog struct {
	ID           uuid.UUID              `json:"id" db:"id"`
	UserID       *uuid.UUID             `json:"user_id,omitempty" db:"user_id"`
	Action       string                 `json:"action" db:"action"`
	ResourceType *string                `json:"resource_type,omitempty" db:"resource_type"`
	ResourceID   *uuid.UUID             `json:"resource_id,omitempty" db:"resource_id"`
	IPAddress    *string                `json:"ip_address,omitempty" db:"ip_address"`
	UserAgent    *string                `json:"user_agent,omitempty" db:"user_agent"`
	Metadata     map[string]interface{} `json:"metadata,omitempty" db:"metadata"`
	CreatedAt    time.Time              `json:"created_at" db:"created_at"`
}
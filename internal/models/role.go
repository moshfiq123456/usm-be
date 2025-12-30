package models

import (
	"time"

	"github.com/google/uuid"
)

// Role represents the roles table
type Role struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Permission represents the permissions table
type Permission struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Resource    string    `json:"resource" db:"resource"`
	Action      string    `json:"action" db:"action"`
	Description *string   `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// UserRole represents the user_roles junction table
type UserRole struct {
	UserID     uuid.UUID `json:"user_id" db:"user_id"`
	RoleID     uuid.UUID `json:"role_id" db:"role_id"`
	AssignedAt time.Time `json:"assigned_at" db:"assigned_at"`
}

// RolePermission represents the role_permissions junction table
type RolePermission struct {
	RoleID       uuid.UUID `json:"role_id" db:"role_id"`
	PermissionID uuid.UUID `json:"permission_id" db:"permission_id"`
	GrantedAt    time.Time `json:"granted_at" db:"granted_at"`
}

// RoleWithPermissions combines Role with its Permissions
type RoleWithPermissions struct {
	Role
	Permissions []Permission `json:"permissions"`
}

// UserWithRoles combines User with their Roles
type UserWithRoles struct {
	User
	Roles []Role `json:"roles"`
}
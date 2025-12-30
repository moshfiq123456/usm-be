package models

import "github.com/google/uuid"

// CreateUserRequest represents the request to create a new user
type CreateUserRequest struct {
	Username  string  `json:"username" validate:"required,min=3,max=50"`
	Email     string  `json:"email" validate:"required,email"`
	Password  string  `json:"password" validate:"required,min=8"`
	FirstName *string `json:"first_name,omitempty" validate:"omitempty,max=100"`
	LastName  *string `json:"last_name,omitempty" validate:"omitempty,max=100"`
}

// UpdateUserRequest represents the request to update user information
type UpdateUserRequest struct {
	Username  *string `json:"username,omitempty" validate:"omitempty,min=3,max=50"`
	Email     *string `json:"email,omitempty" validate:"omitempty,email"`
	FirstName *string `json:"first_name,omitempty" validate:"omitempty,max=100"`
	LastName  *string `json:"last_name,omitempty" validate:"omitempty,max=100"`
	Phone     *string `json:"phone,omitempty" validate:"omitempty,max=20"`
	Bio       *string `json:"bio,omitempty"`
}

// LoginRequest represents the login request
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// ChangePasswordRequest represents the request to change password
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

// ResetPasswordRequest represents the request to reset password
type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

// AssignRoleRequest represents the request to assign a role to a user
type AssignRoleRequest struct {
	RoleID uuid.UUID `json:"role_id" validate:"required"`
}
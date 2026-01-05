// internal/users/dto.go
package users

type CreateUserRequest struct {
	Name     string  `json:"name" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=8"`
	Phone    *string `json:"phone"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=active inactive blocked"`
}

type UserResponse struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Phone  *string `json:"phone"`
	Status string  `json:"status"`
}

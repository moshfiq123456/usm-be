package user_hierarchy

import "github.com/moshfiq123456/ums-be/internal/models"


type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToUserResponse(u models.User) UserResponse {
	return UserResponse{
		ID:    u.ID.String(),
		Name:  u.Name,
		Email: u.Email,
	}
}

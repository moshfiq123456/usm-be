// internal/users/service.go
package users

import (
	"errors"

	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
	"gorm.io/gorm"

	"your_project/internal/domain/models"
)

type Service struct {
	repo *Repository
	db   *gorm.DB
}

func NewService(repo *Repository, db *gorm.DB) *Service {
	return &Service{repo: repo, db: db}
}

func (s *Service) Create(ctx *gofr.Context, req CreateUserRequest) (*UserResponse, error) {
	user := models.User{
		ID:           uuid.New(),
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hashPassword(req.Password),
		Phone:        req.Phone,
		Status:       "active",
	}

	if err := s.repo.Create(ctx, &user); err != nil {
		return nil, err
	}

	return toResponse(&user), nil
}

func (s *Service) List(ctx *gofr.Context) ([]UserResponse, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]UserResponse, 0, len(users))
	for _, u := range users {
		resp = append(resp, *toResponse(&u))
	}

	return resp, nil
}

func (s *Service) GetByID(ctx *gofr.Context, id string) (*UserResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	user, err := s.repo.FindByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return toResponse(user), nil
}

func (s *Service) Update(ctx *gofr.Context, id string, req UpdateUserRequest) (*UserResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	user, err := s.repo.FindByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Phone != nil {
		user.Phone = req.Phone
	}

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return toResponse(user), nil
}

func (s *Service) Delete(ctx *gofr.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid user id")
	}

	return s.repo.Delete(ctx, uid)
}

func (s *Service) UpdateStatus(ctx *gofr.Context, id string, status string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid user id")
	}

	return s.repo.UpdateStatus(ctx, uid, status)
}

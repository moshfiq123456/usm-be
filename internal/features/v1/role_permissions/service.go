package role_permissions

import (
	"context"

	"github.com/moshfiq123456/ums-be/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AssignPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error {
	return s.repo.AssignPermissions(ctx, roleID, permissionIDs)
}

func (s *Service) RemovePermissions(ctx context.Context, roleID uint, permissionIDs []uint) error {
	return s.repo.RemovePermissions(ctx, roleID, permissionIDs)
}

func (s *Service) ListPermissions(ctx context.Context, roleID uint) ([]models.Permission, error) {
	return s.repo.ListPermissions(ctx, roleID)
}

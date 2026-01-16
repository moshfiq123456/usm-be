package user_hierarchy

import (
	"errors"

	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
)

type Service struct {
	repo *Repository
}

func NewService() *Service {
	return &Service{
		repo: NewRepository(),
	}
}

func (s *Service) AssignChild(ctx *gofr.Context, parentID, childID uuid.UUID) (interface{}, error) {
	if parentID == childID {
		return nil, errors.New("parent and child cannot be same")
	}

	if s.repo.Exists(ctx, parentID, childID) {
		return nil, errors.New("hierarchy already exists")
	}

	if s.repo.IsDescendant(ctx, childID, parentID) {
		return nil, errors.New("circular hierarchy detected")
	}

	return s.repo.Create(ctx, parentID, childID)
}

func (s *Service) RemoveChild(ctx *gofr.Context, parentID, childID uuid.UUID) (interface{}, error) {
	return nil, s.repo.Delete(ctx, parentID, childID)
}

func (s *Service) GetChildren(ctx *gofr.Context, userID uuid.UUID) (interface{}, error) {
	return s.repo.GetChildren(ctx, userID)
}

func (s *Service) GetParent(ctx *gofr.Context, userID uuid.UUID) (interface{}, error) {
	return s.repo.GetParent(ctx, userID)
}

func (s *Service) CheckHierarchy(ctx *gofr.Context, parentID, childID uuid.UUID) (interface{}, error) {
	exists := s.repo.IsDescendant(ctx, parentID, childID)
	return CheckHierarchyResponse{IsRelated: exists}, nil
}

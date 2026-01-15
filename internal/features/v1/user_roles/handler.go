package user_roles

import (
	"fmt"

	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) AssignRoles(ctx *gofr.Context) (interface{}, error) {
	userIDStr := ctx.PathParam("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	var req AssignRolesRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	if err := h.service.AssignRoles(ctx, userID, req.RoleIDs); err != nil {
		return nil, err
	}

	return map[string]string{"message": "Roles assigned to user successfully"}, nil
}

func (h *Handler) RemoveRoles(ctx *gofr.Context) (interface{}, error) {
	userIDStr := ctx.PathParam("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	var req RemoveRolesRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	if err := h.service.RemoveRoles(ctx, userID, req.RoleIDs); err != nil {
		return nil, err
	}

	return map[string]string{"message": "Roles removed from user successfully"}, nil
}

func (h *Handler) ListRoles(ctx *gofr.Context) (interface{}, error) {
	userIDStr := ctx.PathParam("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	roles, err := h.service.ListRoles(ctx, userID)
	if err != nil {
		return nil, err
	}

	resp := ToUserRoleResponses(userIDStr, roles)
	return resp, nil
}

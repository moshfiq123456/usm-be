package user_permissions

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

func (h *Handler) AssignPermissions(ctx *gofr.Context) (interface{}, error) {
	userIDStr := ctx.PathParam("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	var req AssignPermissionsRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	if err := h.service.AssignPermissions(ctx, userID, req.PermissionIDs); err != nil {
		return nil, err
	}

	return map[string]string{"message": "Permissions assigned successfully"}, nil
}

func (h *Handler) RemovePermissions(ctx *gofr.Context) (interface{}, error) {
	userIDStr := ctx.PathParam("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	var req RemovePermissionsRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	if err := h.service.RemovePermissions(ctx, userID, req.PermissionIDs); err != nil {
		return nil, err
	}

	return map[string]string{"message": "Permissions removed successfully"}, nil
}

func (h *Handler) ListPermissions(ctx *gofr.Context) (interface{}, error) {
	userIDStr := ctx.PathParam("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	perms, err := h.service.ListPermissions(ctx, userID)
	if err != nil {
		return nil, err
	}

	resp := ToUserPermissionResponses(userIDStr, perms)
	return resp, nil
}

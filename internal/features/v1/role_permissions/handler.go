package role_permissions

import (
	"strconv"

	"gofr.dev/pkg/gofr"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) AssignPermissions(ctx *gofr.Context) (interface{}, error) {
	roleID, err := strconv.ParseUint(ctx.PathParam("id"), 10, 64)
	if err != nil {
		return nil, err
	}

	var req AssignPermissionsRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	if err := h.service.AssignPermissions(ctx, uint(roleID), req.PermissionIDs); err != nil {
		return nil, err
	}

	return map[string]string{"message": "Permissions assigned to role successfully"}, nil
}

func (h *Handler) RemovePermissions(ctx *gofr.Context) (interface{}, error) {
	roleID, err := strconv.ParseUint(ctx.PathParam("id"), 10, 64)
	if err != nil {
		return nil, err
	}

	var req RemovePermissionsRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	if err := h.service.RemovePermissions(ctx, uint(roleID), req.PermissionIDs); err != nil {
		return nil, err
	}

	return map[string]string{"message": "Permissions removed from role successfully"}, nil
}

func (h *Handler) ListPermissions(ctx *gofr.Context) (interface{}, error) {
	roleID, err := strconv.ParseUint(ctx.PathParam("id"), 10, 64)
	if err != nil {
		return nil, err
	}

	perms, err := h.service.ListPermissions(ctx, uint(roleID))
	if err != nil {
		return nil, err
	}

	return ToRolePermissionResponses(uint(roleID), perms), nil
}

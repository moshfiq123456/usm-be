// internal/features/v1/roles/handler.go
package roles

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

func (h *Handler) Create(ctx *gofr.Context) (interface{}, error) {
	var req CreateRoleRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	role, err := h.service.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return toResponse(role), nil
}

func (h *Handler) List(ctx *gofr.Context) (interface{}, error) {
	roles, err := h.service.List(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]RoleResponse, 0, len(roles))
	for _, r := range roles {
		resp = append(resp, toResponse(r))
	}

	return resp, nil
}

func (h *Handler) Get(ctx *gofr.Context) (interface{}, error) {
	id, _ := strconv.ParseInt(ctx.PathParam("id"), 10, 64)
	role, err := h.service.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return toResponse(role), nil
}

func (h *Handler) Update(ctx *gofr.Context) (interface{}, error) {
	id, _ := strconv.ParseInt(ctx.PathParam("id"), 10, 64)

	var req UpdateRoleRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	role, err := h.service.Update(ctx, id, req)
	if err != nil {
		return nil, err
	}

	return toResponse(role), nil
}

func (h *Handler) SetStatus(ctx *gofr.Context) (interface{}, error) {
	id, _ := strconv.ParseInt(ctx.PathParam("id"), 10, 64)

	var req UpdateRoleStatusRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	if err := h.service.UpdateStatus(ctx, id, req.IsActive); err != nil {
		return nil, err
	}

	return map[string]string{"message": "Role status updated successfully"}, nil
}

func (h *Handler) Delete(ctx *gofr.Context) (interface{}, error) {
	id, _ := strconv.ParseInt(ctx.PathParam("id"), 10, 64)

	if err := h.service.Delete(ctx, id); err != nil {
		return nil, err
	}

	return map[string]string{"message": "Role deleted successfully"}, nil
}

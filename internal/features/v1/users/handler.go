package users

import (
	"gofr.dev/pkg/gofr"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// POST /users
func (h *Handler) CreateUser(ctx *gofr.Context) (interface{}, error) {
	var req CreateUserRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	user, err := h.service.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return toResponse(user), nil
}

// GET /users
func (h *Handler) ListUsers(ctx *gofr.Context) (interface{}, error) {
	users, err := h.service.List(ctx)
	if err != nil {
		return nil, err
	}

	return toResponseList(users), nil
}

// GET /users/{id}
func (h *Handler) GetUser(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	user, err := h.service.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return toResponse(user), nil
}

// PUT /users/{id}
func (h *Handler) UpdateUser(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	var req UpdateUserRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	user, err := h.service.Update(ctx, id, req)
	if err != nil {
		return nil, err
	}

	return toResponse(user), nil
}

// DELETE /users/{id}
func (h *Handler) DeleteUser(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	if err := h.service.Delete(ctx, id); err != nil {
		return nil, err
	}

	return map[string]string{"message": "user deleted"}, nil
}

// PATCH /users/{id}/status
func (h *Handler) SetStatus(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	var req UpdateStatusRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	if err := h.service.UpdateStatus(ctx, id, req.Status); err != nil {
		return nil, err
	}

	return map[string]string{"message": "status updated"}, nil
}

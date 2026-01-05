// internal/users/handler.go
package users

import (
	"github.com/moshfiq123456/ums-be/internal/pkg/validator"
	"gofr.dev/pkg/gofr"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateUser(ctx *gofr.Context) (interface{}, error) {
	var req CreateUserRequest

	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	if err := validator.Validate.Struct(req); err != nil {
		return nil, err
	}

	return h.service.Create(ctx, req)
}


func (h *Handler) ListUsers(ctx *gofr.Context) (interface{}, error) {
	return h.service.List(ctx)
}

func (h *Handler) GetUser(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	return h.service.GetByID(ctx, id)
}

func (h *Handler) UpdateUser(ctx *gofr.Context) (interface{}, error) {
	var req UpdateUserRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	if err := validator.Validate.Struct(req); err != nil {
		return nil, err
	}

	return h.service.Update(ctx, ctx.PathParam("id"), req)
}


func (h *Handler) DeleteUser(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	return nil, h.service.Delete(ctx, id)
}

func (h *Handler) SetStatus(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	var req UpdateStatusRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	return nil, h.service.UpdateStatus(ctx, id, req.Status)
}

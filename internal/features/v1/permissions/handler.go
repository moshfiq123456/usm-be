package permissions

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
	var req CreatePermissionRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	err := h.service.Create(ctx, req)
	return map[string]string{"message": "permission created"}, err
}

func (h *Handler) List(ctx *gofr.Context) (interface{}, error) {
	perms, err := h.service.List(ctx)
	if err != nil {
		return nil, err
	}
	return ToResponseList(perms), nil
}


func (h *Handler) Get(ctx *gofr.Context) (interface{}, error) {
	id, _ := strconv.Atoi(ctx.PathParam("id"))

	p, err := h.service.Get(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	return ToResponse(p), nil
}


func (h *Handler) Update(ctx *gofr.Context) (interface{}, error) {
	id, _ := strconv.Atoi(ctx.PathParam("id"))

	var req UpdatePermissionRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	return h.service.Update(ctx, uint(id), req)
}

func (h *Handler) Delete(ctx *gofr.Context) (interface{}, error) {
	id, _ := strconv.Atoi(ctx.PathParam("id"))
	return map[string]string{"message": "permission deleted"}, h.service.Delete(ctx, uint(id))
}

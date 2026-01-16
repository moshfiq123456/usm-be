package user_hierarchy

import (
	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
)

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService(),
	}
}

func (h *Handler) AssignChild(ctx *gofr.Context) (interface{}, error) {
	parentID := uuid.MustParse(ctx.PathParam("parentId"))

	var req AssignChildRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	return h.service.AssignChild(ctx, parentID, req.ChildUserID)
}

func (h *Handler) RemoveChild(ctx *gofr.Context) (interface{}, error) {
	parentID := uuid.MustParse(ctx.PathParam("parentId"))
	childID := uuid.MustParse(ctx.PathParam("childId"))

	return h.service.RemoveChild(ctx, parentID, childID)
}

func (h *Handler) GetChildren(ctx *gofr.Context) (interface{}, error) {
	userID := uuid.MustParse(ctx.PathParam("id"))
	return h.service.GetChildren(ctx, userID)
}

func (h *Handler) GetParent(ctx *gofr.Context) (interface{}, error) {
	userID := uuid.MustParse(ctx.PathParam("id"))
	return h.service.GetParent(ctx, userID)
}

func (h *Handler) CheckHierarchy(ctx *gofr.Context) (interface{}, error) {
	var req CheckHierarchyRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	return h.service.CheckHierarchy(ctx, req.ParentID, req.ChildID)
}

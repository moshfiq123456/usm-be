package user_hierarchy

import "github.com/google/uuid"

type AssignChildRequest struct {
	ChildUserID uuid.UUID `json:"child_user_id" validate:"required"`
}

type CheckHierarchyRequest struct {
	ParentID uuid.UUID `json:"parent_id" validate:"required"`
	ChildID  uuid.UUID `json:"child_id" validate:"required"`
}

type CheckHierarchyResponse struct {
	IsRelated bool `json:"is_related"`
}

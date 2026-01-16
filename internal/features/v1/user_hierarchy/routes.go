package user_hierarchy

import "github.com/gofr-dev/gofr/pkg/gofr"

func registerRoutes(app *gofr.App) {
	h := NewHandler()

	app.POST("/v1/users/{parentId}/children", h.AssignChild)
	app.DELETE("/v1/users/{parentId}/children/{childId}", h.RemoveChild)

	app.GET("/v1/users/{id}/children", h.GetChildren)
	app.GET("/v1/users/{id}/parent", h.GetParent)

	app.POST("/v1/users/hierarchy/check", h.CheckHierarchy)
}

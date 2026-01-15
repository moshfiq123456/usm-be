package user_permissions

import "gofr.dev/pkg/gofr"

func RegisterRoutes(app *gofr.App, handler *Handler) {
	app.POST("/users/{id}/permissions", handler.AssignPermissions)
	app.DELETE("/users/{id}/permissions", handler.RemovePermissions)
	app.GET("/users/{id}/permissions", handler.ListPermissions)
}

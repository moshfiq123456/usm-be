package role_permissions

import "gofr.dev/pkg/gofr"

func RegisterRoutes(app *gofr.App, handler *Handler) {
	app.POST("/roles/{id}/permissions", handler.AssignPermissions)
	app.DELETE("/roles/{id}/permissions", handler.RemovePermissions)
	app.GET("/roles/{id}/permissions", handler.ListPermissions)
}

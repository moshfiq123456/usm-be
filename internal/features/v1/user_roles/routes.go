package user_roles

import "gofr.dev/pkg/gofr"

func RegisterRoutes(app *gofr.App, handler *Handler) {
	app.POST("/users/{id}/roles", handler.AssignRoles)
	app.DELETE("/users/{id}/roles", handler.RemoveRoles)
	app.GET("/users/{id}/roles", handler.ListRoles)
}

// internal/features/v1/roles/routes.go
package roles

import "gofr.dev/pkg/gofr"

func RegisterRoutes(app *gofr.App, handler *Handler) {
	app.POST("/roles", handler.Create)
	app.GET("/roles", handler.List)
	app.GET("/roles/{id}", handler.Get)
	app.PUT("/roles/{id}", handler.Update)
	app.PATCH("/roles/{id}/status", handler.SetStatus)
	app.DELETE("/roles/{id}", handler.Delete)
}

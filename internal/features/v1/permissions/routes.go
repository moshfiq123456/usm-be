package permissions

import "gofr.dev/pkg/gofr"

func RegisterRoutes(app *gofr.App, handler *Handler) {
	app.POST("/permissions", handler.Create)
	app.GET("/permissions", handler.List)
	app.GET("/permissions/{id}", handler.Get)
	app.PUT("/permissions/{id}", handler.Update)
	app.DELETE("/permissions/{id}", handler.Delete)
}

package users

import "gofr.dev/pkg/gofr"

func RegisterRoutes(app *gofr.App, handler *Handler) {

	app.POST("/users", handler.CreateUser)
	app.GET("/users", handler.ListUsers)
	app.GET("/users/{id}", handler.GetUser)
	app.PUT("/users/{id}", handler.UpdateUser)
	app.DELETE("/users/{id}", handler.DeleteUser)
	app.PATCH("/users/{id}/status", handler.SetStatus)
}

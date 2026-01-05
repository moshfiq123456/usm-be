// internal/users/routes.go
package users

import "gofr.dev/pkg/gofr"

func RegisterRoutes(app *gofr.App, handler *Handler) {
	users := app.Group("/users")

	users.POST("", handler.CreateUser)           // POST /users
	users.GET("", handler.ListUsers)              // GET /users
	users.GET("/{id}", handler.GetUser)           // GET /users/{id}
	users.PUT("/{id}", handler.UpdateUser)        // PUT /users/{id}
	users.DELETE("/{id}", handler.DeleteUser)     // DELETE /users/{id}
	users.PATCH("/{id}/status", handler.SetStatus) // PATCH /users/{id}/status
}

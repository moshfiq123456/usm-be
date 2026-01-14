package users

import (
	"gofr.dev/pkg/gofr"
	"gorm.io/gorm"
)

func RegisterModule(app *gofr.App, db *gorm.DB) {
	repo := NewUserRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	RegisterRoutes(app, handler)
}

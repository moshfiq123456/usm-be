// internal/features/v1/roles/module.go
package roles

import (
	"gofr.dev/pkg/gofr"
	"gorm.io/gorm"
)

func RegisterModule(app *gofr.App, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	RegisterRoutes(app, handler)
}

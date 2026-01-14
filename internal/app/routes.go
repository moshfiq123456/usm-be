package app

import (
	"gofr.dev/pkg/gofr"
	"gorm.io/gorm"

	"github.com/moshfiq123456/ums-be/internal/features/v1/users"
)

func RegisterRoutes(app *gofr.App, db *gorm.DB) {
	users.RegisterModule(app, db)
}

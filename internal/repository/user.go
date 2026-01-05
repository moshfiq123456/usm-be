// internal/users/repository.go
package users

import (
	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
	"gorm.io/gorm"

	"your_project/internal/domain/models"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx *gofr.Context, user *models.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) FindAll(ctx *gofr.Context) ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *Repository) FindByID(ctx *gofr.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	return &user, err
}

func (r *Repository) Update(ctx *gofr.Context, user *models.User) error {
	return r.db.Save(user).Error
}

func (r *Repository) Delete(ctx *gofr.Context, id uuid.UUID) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}

func (r *Repository) UpdateStatus(ctx *gofr.Context, id uuid.UUID, status string) error {
	return r.db.Model(&models.User{}).
		Where("id = ?", id).
		Update("status", status).Error
}

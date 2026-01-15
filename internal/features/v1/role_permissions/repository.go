package role_permissions

import (
	"context"

	"github.com/moshfiq123456/ums-be/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AssignPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error {
	for _, pid := range permissionIDs {
		if err := r.db.WithContext(ctx).FirstOrCreate(
			&models.RolePermission{},
			models.RolePermission{
				RoleID:       roleID,
				PermissionID: pid,
			},
		).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) RemovePermissions(ctx context.Context, roleID uint, permissionIDs []uint) error {
	return r.db.WithContext(ctx).
		Where("role_id = ? AND permission_id IN ?", roleID, permissionIDs).
		Delete(&models.RolePermission{}).Error
}

func (r *Repository) ListPermissions(ctx context.Context, roleID uint) ([]models.Permission, error) {
	var permissions []models.Permission

	err := r.db.WithContext(ctx).
		Model(&models.Permission{}).
		Joins("JOIN role_permissions rp ON rp.permission_id = permissions.id").
		Where("rp.role_id = ?", roleID).
		Find(&permissions).Error

	return permissions, err
}

package user_hierarchy

import (
	"github.com/gofr-dev/gofr/pkg/gofr"
	"github.com/google/uuid"
	"github.com/moshfiq123456/ums-be/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) db(ctx *gofr.Context) *gorm.DB {
	return ctx.DB()
}

func (r *Repository) Exists(ctx *gofr.Context, parentID, childID uuid.UUID) bool {
	var count int64
	r.db(ctx).
		Model(&models.UserHierarchy{}).
		Where("parent_user_id = ? AND child_user_id = ?", parentID, childID).
		Count(&count)
	return count > 0
}

func (r *Repository) Create(ctx *gofr.Context, parentID, childID uuid.UUID) error {
	return r.db(ctx).Create(&models.UserHierarchy{
		ParentUserID: parentID,
		ChildUserID:  childID,
	}).Error
}

func (r *Repository) Delete(ctx *gofr.Context, parentID, childID uuid.UUID) error {
	return r.db(ctx).
		Where("parent_user_id = ? AND child_user_id = ?", parentID, childID).
		Delete(&models.UserHierarchy{}).Error
}

func (r *Repository) GetChildren(ctx *gofr.Context, userID uuid.UUID) (interface{}, error) {
	var users []models.User
	err := r.db(ctx).
		Joins("JOIN user_hierarchies uh ON uh.child_user_id = users.id").
		Where("uh.parent_user_id = ?", userID).
		Find(&users).Error
	return users, err
}

func (r *Repository) GetParent(ctx *gofr.Context, userID uuid.UUID) (interface{}, error) {
	var user models.User
	err := r.db(ctx).
		Joins("JOIN user_hierarchies uh ON uh.parent_user_id = users.id").
		Where("uh.child_user_id = ?", userID).
		First(&user).Error
	return user, err
}

/*
OPTIONAL:
func (r *Repository) IsDescendant(ctx *gofr.Context, parentID, childID uuid.UUID) bool {
    // Recursive CTE implementation if needed
}
*/

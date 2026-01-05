package models

type Permission struct {
	ID          uint   `gorm:"primaryKey"`
	Code        string `gorm:"size:100;uniqueIndex;not null"` // e.g. user.create
	Description string `gorm:"size:255"`

	Roles []RolePermission `gorm:"foreignKey:PermissionID"`
	Users []UserPermission `gorm:"foreignKey:PermissionID"`
}

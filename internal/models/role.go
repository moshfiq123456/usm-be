package models

type Role struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:50;uniqueIndex;not null"`
	Description string `gorm:"size:255"`

	Users       []User           `gorm:"many2many:user_roles"`
	Permissions []RolePermission `gorm:"foreignKey:RoleID"`
}

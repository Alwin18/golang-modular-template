package models

type UserRole struct {
	UserID uint `gorm:"not null;primaryKey;index:idx_user_roles_user_id"`
	RoleID uint `gorm:"not null;primaryKey;index:idx_user_roles_role_id"`

	// Relationships
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL;"`
	Role *Role `gorm:"foreignKey:RoleID;constraint:OnDelete:SET NULL;"`
}

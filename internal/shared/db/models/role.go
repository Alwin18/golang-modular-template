package models

type Role struct {
	ID          uint    `gorm:"primary_key" json:"id"`
	Code        string  `gorm:"type:varchar(50);not null;uniqueIndex:ux_roles_code" json:"code"`
	Name        string  `gorm:"type:varchar(255);not null" json:"name"`
	Description *string `gorm:"type:text" json:"description"`

	// Relationships
	Users []User `gorm:"many2many:user_roles;constraint:OnDelete:SET NULL;" json:"users,omitempty"`
}

// TableName overrides the table name
func (Role) TableName() string {
	return "roles"
}

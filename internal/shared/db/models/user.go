package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null;uniqueIndex:ux_users_username" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null;column:password_hash" json:"password_hash"`
	FullName  string    `gorm:"type:varchar(255);not null;column:full_name" json:"full_name"`
	Email     *string   `gorm:"type:varchar(255);uniqueIndex:ux_users_email" json:"email"`
	Phone     *string   `gorm:"type:varchar(50)" json:"phone"`
	IsActive  *bool     `gorm:"type:boolean;not null;default:true;column:is_active" json:"is_active"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	Roles []Role `gorm:"many2many:user_roles;constraint:OnDelete:SET NULL;" json:"roles,omitempty"`
}

func (User) TableName() string {
	return "users"
}

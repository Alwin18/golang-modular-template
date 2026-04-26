package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null" json:"username"`
	Email    string `gorm:"type:varchar(150);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Status   string `gorm:"type:varchar(50);default:'ACTIVE'" json:"status"`
	RoleID   uint   `gorm:"not null"`

	// Relationship
	Role *Role `gorm:"foreignKey:RoleID" json:"role,omitempty"`
}

func (u *User) TableName() string {
	return "users"
}

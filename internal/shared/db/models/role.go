package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Status      string `gorm:"type:varchar(50);default:'ACTIVE'" json:"status"`
}

func (r *Role) TableName() string {
	return "roles"
}

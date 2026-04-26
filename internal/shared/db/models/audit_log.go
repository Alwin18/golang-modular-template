package models

import "time"

type AuditLog struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     *uint     `gorm:"index" json:"user_id"`
	Action     string    `gorm:"type:varchar(100);not null" json:"action"`
	Resource   string    `gorm:"type:varchar(100)" json:"resource"`
	ResourceID *uint     `json:"resource_id"`
	IPAddress  string    `gorm:"type:varchar(50)" json:"ip_address"`
	UserAgent  string    `gorm:"type:varchar(500)" json:"user_agent"`
	CreatedAt  time.Time `json:"created_at"`
}

package models

import "time"

type StockCount struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CountNumber string     `gorm:"type:varchar(100);not null;uniqueIndex:ux_stockcount_number;column:count_number" json:"count_number"`
	WarehouseID uint       `gorm:"not null;index:idx_sc_wh" json:"warehouse_id"`
	Status      string     `gorm:"type:varchar(50);not null" json:"status"`
	CountType   string     `gorm:"type:varchar(50);not null;column:count_type" json:"count_type"`
	ScheduledAt *time.Time `gorm:"column:scheduled_at" json:"scheduled_at"`
	CompletedAt *time.Time `gorm:"column:completed_at" json:"completed_at"`
	CreatedBy   uint       `gorm:"not null;index:idx_sc_created_by" json:"created_by"`
	CreatedAt   time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`

	// Relationships
	Warehouse *Warehouse       `gorm:"foreignKey:WarehouseID;constraint:OnDelete:RESTRICT;" json:"warehouse,omitempty"`
	Creator   *User            `gorm:"foreignKey:CreatedBy;constraint:OnDelete:RESTRICT;" json:"creator,omitempty"`
	Lines     []StockCountLine `gorm:"foreignKey:StockCountID" json:"lines,omitempty"`
}

func (StockCount) TableName() string {
	return "stock_counts"
}

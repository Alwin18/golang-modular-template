package models

import "time"

type StockAdjustment struct {
	ID               uint       `gorm:"primary_key" json:"id"`
	AdjustmentNumber string     `gorm:"type:varchar(100);not null;uniqueIndex:ux_adj_number;column:adjustment_number" json:"adjustment_number"`
	WarehouseID      uint       `gorm:"not null;index:idx_adj_wh" json:"warehouse_id"`
	ReasonCode       string     `gorm:"type:varchar(50);not null;column:reason_code" json:"reason_code"`
	Status           string     `gorm:"type:varchar(50);not null" json:"status"`
	CreatedBy        uint       `gorm:"not null;index:idx_adj_created_by" json:"created_by"`
	CreatedAt        time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	PostedAt         *time.Time `gorm:"column:posted_at" json:"posted_at"`

	// Relationships
	Warehouse *Warehouse            `gorm:"foreignKey:WarehouseID;constraint:OnDelete:RESTRICT;" json:"warehouse,omitempty"`
	Creator   *User                 `gorm:"foreignKey:CreatedBy;constraint:OnDelete:RESTRICT;" json:"creator,omitempty"`
	Lines     []StockAdjustmentLine `gorm:"foreignKey:StockAdjustmentID" json:"lines,omitempty"`
}

func (StockAdjustment) TableName() string {
	return "stock_adjustments"
}

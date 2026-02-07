package models

import "time"

type SupplierReturn struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	ReturnNumber string    `gorm:"type:varchar(100);not null;uniqueIndex:ux_supret_number;column:return_number" json:"return_number"`
	SupplierID   uint      `gorm:"not null;index:idx_supret_supplier" json:"supplier_id"`
	WarehouseID  uint      `gorm:"not null;index:idx_supret_wh" json:"warehouse_id"`
	Status       string    `gorm:"type:varchar(50);not null" json:"status"`
	Reason       *string   `gorm:"type:text" json:"reason"`
	CreatedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy    uint      `gorm:"not null;index:idx_supret_created_by" json:"created_by"`

	// Relationships
	Supplier  *Supplier            `gorm:"foreignKey:SupplierID;constraint:OnDelete:RESTRICT;" json:"supplier,omitempty"`
	Warehouse *Warehouse           `gorm:"foreignKey:WarehouseID;constraint:OnDelete:RESTRICT;" json:"warehouse,omitempty"`
	Creator   *User                `gorm:"foreignKey:CreatedBy;constraint:OnDelete:RESTRICT;" json:"creator,omitempty"`
	Lines     []SupplierReturnLine `gorm:"foreignKey:SupplierReturnID" json:"lines,omitempty"`
}

func (SupplierReturn) TableName() string {
	return "supplier_returns"
}

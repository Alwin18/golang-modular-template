package models

import "time"

type InventoryMovement struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	MovementType   string    `gorm:"type:varchar(50);not null;column:movement_type" json:"movement_type"`
	WarehouseID    uint      `gorm:"not null;index:idx_invmov_wh" json:"warehouse_id"`
	ProductID      uint      `gorm:"not null;index:idx_invmov_product" json:"product_id"`
	BatchID        *uint     `gorm:"index:idx_invmov_batch" json:"batch_id"`
	FromLocationID *uint     `gorm:"index:idx_invmov_fromloc" json:"from_location_id"`
	ToLocationID   *uint     `gorm:"index:idx_invmov_toloc" json:"to_location_id"`
	Qty            float64   `gorm:"type:decimal(18,3);not null" json:"qty"`
	UOM            string    `gorm:"type:varchar(20);not null" json:"uom"`
	StatusBefore   *string   `gorm:"type:varchar(50);column:status_before" json:"status_before"`
	StatusAfter    *string   `gorm:"type:varchar(50);column:status_after" json:"status_after"`
	ReferenceType  *string   `gorm:"type:varchar(50);column:reference_type" json:"reference_type"`
	ReferenceID    *uint     `gorm:"column:reference_id" json:"reference_id"`
	CreatedAt      time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;index:idx_invmov_created" json:"created_at"`
	CreatedBy      *uint     `gorm:"index:idx_invmov_created_by" json:"created_by"`
	Note           *string   `gorm:"type:text" json:"note"`

	// Relationships
	Warehouse    *Warehouse    `gorm:"foreignKey:WarehouseID;constraint:OnDelete:CASCADE;" json:"warehouse,omitempty"`
	Product      *Product      `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;" json:"product,omitempty"`
	Batch        *ProductBatch `gorm:"foreignKey:BatchID;constraint:OnDelete:SET NULL;" json:"batch,omitempty"`
	FromLocation *Location     `gorm:"foreignKey:FromLocationID;constraint:OnDelete:SET NULL;" json:"from_location,omitempty"`
	ToLocation   *Location     `gorm:"foreignKey:ToLocationID;constraint:OnDelete:SET NULL;" json:"to_location,omitempty"`
	Creator      *User         `gorm:"foreignKey:CreatedBy;constraint:OnDelete:SET NULL;" json:"creator,omitempty"`
}

func (InventoryMovement) TableName() string {
	return "inventory_movements"
}

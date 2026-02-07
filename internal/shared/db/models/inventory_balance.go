package models

import "time"

type InventoryBalance struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	WarehouseID  uint      `gorm:"not null;index:idx_invbal_wh" json:"warehouse_id"`
	LocationID   uint      `gorm:"not null;index:idx_invbal_loc" json:"location_id"`
	ProductID    uint      `gorm:"not null;index:idx_invbal_product" json:"product_id"`
	BatchID      *uint     `gorm:"index:idx_invbal_batch" json:"batch_id"`
	Status       string    `gorm:"type:varchar(50);not null" json:"status"`
	OnHandQty    float64   `gorm:"type:decimal(18,3);not null;default:0;column:on_hand_qty" json:"on_hand_qty"`
	ReservedQty  float64   `gorm:"type:decimal(18,3);not null;default:0;column:reserved_qty" json:"reserved_qty"`
	AvailableQty float64   `gorm:"type:decimal(18,3);not null;default:0;column:available_qty" json:"available_qty"`
	UpdatedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	Warehouse *Warehouse    `gorm:"foreignKey:WarehouseID;constraint:OnDelete:CASCADE;" json:"warehouse,omitempty"`
	Location  *Location     `gorm:"foreignKey:LocationID;constraint:OnDelete:CASCADE;" json:"location,omitempty"`
	Product   *Product      `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;" json:"product,omitempty"`
	Batch     *ProductBatch `gorm:"foreignKey:BatchID;constraint:OnDelete:SET NULL;" json:"batch,omitempty"`
}

func (InventoryBalance) TableName() string {
	return "inventory_balances"
}

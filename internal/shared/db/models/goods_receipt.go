package models

import "time"

type GoodsReceipt struct {
	ID              uint      `gorm:"primary_key" json:"id"`
	GRNumber        string    `gorm:"type:varchar(100);not null;uniqueIndex:ux_gr_number;column:gr_number" json:"gr_number"`
	PurchaseOrderID *uint     `gorm:"index:idx_gr_po" json:"purchase_order_id"`
	WarehouseID     uint      `gorm:"not null;index:idx_gr_warehouse" json:"warehouse_id"`
	SupplierID      *uint     `gorm:"index:idx_gr_supplier" json:"supplier_id"`
	Status          string    `gorm:"type:varchar(50);not null" json:"status"`
	ReceivedAt      time.Time `gorm:"not null;column:received_at" json:"received_at"`
	ReceivedBy      uint      `gorm:"not null;index:idx_gr_received_by" json:"received_by"`
	ExternalRef     *string   `gorm:"type:varchar(255);column:external_ref" json:"external_ref"`

	// Relationships
	PurchaseOrder *PurchaseOrder     `gorm:"foreignKey:PurchaseOrderID;constraint:OnDelete:SET NULL;" json:"purchase_order,omitempty"`
	Warehouse     *Warehouse         `gorm:"foreignKey:WarehouseID;constraint:OnDelete:RESTRICT;" json:"warehouse,omitempty"`
	Supplier      *Supplier          `gorm:"foreignKey:SupplierID;constraint:OnDelete:SET NULL;" json:"supplier,omitempty"`
	Receiver      *User              `gorm:"foreignKey:ReceivedBy;constraint:OnDelete:RESTRICT;" json:"receiver,omitempty"`
	Lines         []GoodsReceiptLine `gorm:"foreignKey:GoodsReceiptID" json:"lines,omitempty"`
	PutawayTasks  []PutawayTask      `gorm:"foreignKey:GoodsReceiptID" json:"putaway_tasks,omitempty"`
}

func (GoodsReceipt) TableName() string {
	return "goods_receipts"
}

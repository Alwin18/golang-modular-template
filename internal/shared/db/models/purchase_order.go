package models

import "time"

type PurchaseOrder struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	PONumber     string     `gorm:"type:varchar(100);not null;uniqueIndex:ux_po_number;column:po_number" json:"po_number"`
	SupplierID   uint       `gorm:"not null;index:idx_po_supplier" json:"supplier_id"`
	WarehouseID  uint       `gorm:"not null;index:idx_po_warehouse" json:"warehouse_id"`
	Status       string     `gorm:"type:varchar(50);not null" json:"status"`
	OrderDate    time.Time  `gorm:"type:date;not null;column:order_date" json:"order_date"`
	ExpectedDate *time.Time `gorm:"type:date;column:expected_date" json:"expected_date"`
	Currency     *string    `gorm:"type:varchar(10)" json:"currency"`
	TotalAmount  *float64   `gorm:"type:decimal(18,2);column:total_amount" json:"total_amount"`
	CreatedBy    uint       `gorm:"not null;index:idx_po_created_by" json:"created_by"`
	CreatedAt    time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	Supplier      *Supplier           `gorm:"foreignKey:SupplierID;constraint:OnDelete:RESTRICT;" json:"supplier,omitempty"`
	Warehouse     *Warehouse          `gorm:"foreignKey:WarehouseID;constraint:OnDelete:RESTRICT;" json:"warehouse,omitempty"`
	Creator       *User               `gorm:"foreignKey:CreatedBy;constraint:OnDelete:RESTRICT;" json:"creator,omitempty"`
	Lines         []PurchaseOrderLine `gorm:"foreignKey:PurchaseOrderID" json:"lines,omitempty"`
	GoodsReceipts []GoodsReceipt      `gorm:"foreignKey:PurchaseOrderID" json:"goods_receipts,omitempty"`
}

func (PurchaseOrder) TableName() string {
	return "purchase_orders"
}

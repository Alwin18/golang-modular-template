package models

type PurchaseOrderLine struct {
	ID              uint     `gorm:"primary_key" json:"id"`
	PurchaseOrderID uint     `gorm:"not null;index:idx_pol_po" json:"purchase_order_id"`
	LineNo          int      `gorm:"not null;column:line_no" json:"line_no"`
	ProductID       uint     `gorm:"not null;index:idx_pol_product" json:"product_id"`
	UOM             string   `gorm:"type:varchar(20);not null" json:"uom"`
	OrderedQty      float64  `gorm:"type:decimal(18,3);not null;column:ordered_qty" json:"ordered_qty"`
	ReceivedQty     float64  `gorm:"type:decimal(18,3);not null;default:0;column:received_qty" json:"received_qty"`
	UnitPrice       *float64 `gorm:"type:decimal(18,2);column:unit_price" json:"unit_price"`
	TaxPercent      *float64 `gorm:"type:decimal(5,2);column:tax_percent" json:"tax_percent"`

	// Relationships
	PurchaseOrder     *PurchaseOrder     `gorm:"foreignKey:PurchaseOrderID;constraint:OnDelete:CASCADE;" json:"purchase_order,omitempty"`
	Product           *Product           `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT;" json:"product,omitempty"`
	GoodsReceiptLines []GoodsReceiptLine `gorm:"foreignKey:PurchaseOrderLineID" json:"goods_receipt_lines,omitempty"`
}

func (PurchaseOrderLine) TableName() string {
	return "purchase_order_lines"
}

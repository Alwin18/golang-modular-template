package models

type GoodsReceiptLine struct {
	ID                  uint    `gorm:"primary_key" json:"id"`
	GoodsReceiptID      uint    `gorm:"not null;index:idx_grl_gr" json:"goods_receipt_id"`
	PurchaseOrderLineID *uint   `gorm:"index:idx_grl_pol" json:"purchase_order_line_id"`
	LineNo              int     `gorm:"not null;column:line_no" json:"line_no"`
	ProductID           uint    `gorm:"not null;index:idx_grl_product" json:"product_id"`
	UOM                 string  `gorm:"type:varchar(20);not null" json:"uom"`
	ReceivedQty         float64 `gorm:"type:decimal(18,3);not null;column:received_qty" json:"received_qty"`
	BatchID             *uint   `gorm:"index:idx_grl_batch" json:"batch_id"`
	SerialNumber        *string `gorm:"type:varchar(100);column:serial_number" json:"serial_number"`
	QCStatus            *string `gorm:"type:varchar(50);column:qc_status" json:"qc_status"`
	SourceLocationID    *uint   `gorm:"index:idx_grl_sourceloc" json:"source_location_id"`
	Note                *string `gorm:"type:text" json:"note"`

	// Relationships
	GoodsReceipt      *GoodsReceipt      `gorm:"foreignKey:GoodsReceiptID;constraint:OnDelete:CASCADE;" json:"goods_receipt,omitempty"`
	PurchaseOrderLine *PurchaseOrderLine `gorm:"foreignKey:PurchaseOrderLineID;constraint:OnDelete:SET NULL;" json:"purchase_order_line,omitempty"`
	Product           *Product           `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT;" json:"product,omitempty"`
	Batch             *ProductBatch      `gorm:"foreignKey:BatchID;constraint:OnDelete:SET NULL;" json:"batch,omitempty"`
	SourceLocation    *Location          `gorm:"foreignKey:SourceLocationID;constraint:OnDelete:SET NULL;" json:"source_location,omitempty"`
	PutawayTaskLines  []PutawayTaskLine  `gorm:"foreignKey:GoodsReceiptLineID" json:"putaway_task_lines,omitempty"`
}

func (GoodsReceiptLine) TableName() string {
	return "goods_receipt_lines"
}

package models

type SupplierReturnLine struct {
	ID               uint    `gorm:"primary_key" json:"id"`
	SupplierReturnID uint    `gorm:"not null;index:idx_srlines_sr" json:"supplier_return_id"`
	ProductID        uint    `gorm:"not null;index:idx_srlines_product" json:"product_id"`
	BatchID          *uint   `gorm:"index:idx_srlines_batch" json:"batch_id"`
	UOM              string  `gorm:"type:varchar(20);not null" json:"uom"`
	Qty              float64 `gorm:"type:decimal(18,3);not null" json:"qty"`
	ReasonCode       *string `gorm:"type:varchar(50);column:reason_code" json:"reason_code"`

	// Relationships
	SupplierReturn *SupplierReturn `gorm:"foreignKey:SupplierReturnID;constraint:OnDelete:CASCADE;" json:"supplier_return,omitempty"`
	Product        *Product        `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT;" json:"product,omitempty"`
	Batch          *ProductBatch   `gorm:"foreignKey:BatchID;constraint:OnDelete:SET NULL;" json:"batch,omitempty"`
}

func (SupplierReturnLine) TableName() string {
	return "supplier_return_lines"
}

package models

type StockAdjustmentLine struct {
	ID                uint    `gorm:"primary_key" json:"id"`
	StockAdjustmentID uint    `gorm:"not null;index:idx_adjlines_adj" json:"stock_adjustment_id"`
	LocationID        uint    `gorm:"not null;index:idx_adjlines_loc" json:"location_id"`
	ProductID         uint    `gorm:"not null;index:idx_adjlines_product" json:"product_id"`
	BatchID           *uint   `gorm:"index:idx_adjlines_batch" json:"batch_id"`
	QtyDelta          float64 `gorm:"type:decimal(18,3);not null;column:qty_delta" json:"qty_delta"`
	UOM               string  `gorm:"type:varchar(20);not null" json:"uom"`
	Note              *string `gorm:"type:text" json:"note"`

	// Relationships
	StockAdjustment *StockAdjustment `gorm:"foreignKey:StockAdjustmentID;constraint:OnDelete:CASCADE;" json:"stock_adjustment,omitempty"`
	Location        *Location        `gorm:"foreignKey:LocationID;constraint:OnDelete:RESTRICT;" json:"location,omitempty"`
	Product         *Product         `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT;" json:"product,omitempty"`
	Batch           *ProductBatch    `gorm:"foreignKey:BatchID;constraint:OnDelete:SET NULL;" json:"batch,omitempty"`
}

func (StockAdjustmentLine) TableName() string {
	return "stock_adjustment_lines"
}

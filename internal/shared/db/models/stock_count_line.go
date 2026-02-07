package models

type StockCountLine struct {
	ID           uint     `gorm:"primary_key" json:"id"`
	StockCountID uint     `gorm:"not null;index:idx_scl_sc" json:"stock_count_id"`
	LocationID   uint     `gorm:"not null;index:idx_scl_loc" json:"location_id"`
	ProductID    uint     `gorm:"not null;index:idx_scl_product" json:"product_id"`
	BatchID      *uint    `gorm:"index:idx_scl_batch" json:"batch_id"`
	SystemQty    float64  `gorm:"type:decimal(18,3);not null;column:system_qty" json:"system_qty"`
	CountedQty   *float64 `gorm:"type:decimal(18,3);column:counted_qty" json:"counted_qty"`
	VarianceQty  *float64 `gorm:"type:decimal(18,3);column:variance_qty" json:"variance_qty"`

	// Relationships
	StockCount *StockCount   `gorm:"foreignKey:StockCountID;constraint:OnDelete:CASCADE;" json:"stock_count,omitempty"`
	Location   *Location     `gorm:"foreignKey:LocationID;constraint:OnDelete:RESTRICT;" json:"location,omitempty"`
	Product    *Product      `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT;" json:"product,omitempty"`
	Batch      *ProductBatch `gorm:"foreignKey:BatchID;constraint:OnDelete:SET NULL;" json:"batch,omitempty"`
}

func (StockCountLine) TableName() string {
	return "stock_count_lines"
}

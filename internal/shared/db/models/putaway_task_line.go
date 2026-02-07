package models

type PutawayTaskLine struct {
	ID                    uint    `gorm:"primary_key" json:"id"`
	PutawayTaskID         uint    `gorm:"not null;index:idx_putlines_task" json:"putaway_task_id"`
	GoodsReceiptLineID    uint    `gorm:"not null;index:idx_putlines_grl" json:"goods_receipt_line_id"`
	ProductID             uint    `gorm:"not null;index:idx_putlines_product" json:"product_id"`
	SourceLocationID      uint    `gorm:"not null;index:idx_putlines_sourceloc" json:"source_location_id"`
	DestinationLocationID uint    `gorm:"not null;index:idx_putlines_destloc" json:"destination_location_id"`
	BatchID               *uint   `gorm:"index:idx_putlines_batch" json:"batch_id"`
	UOM                   string  `gorm:"type:varchar(20);not null" json:"uom"`
	PlannedQty            float64 `gorm:"type:decimal(18,3);not null;column:planned_qty" json:"planned_qty"`
	PutawayQty            float64 `gorm:"type:decimal(18,3);not null;default:0;column:putaway_qty" json:"putaway_qty"`

	// Relationships
	PutawayTask         *PutawayTask      `gorm:"foreignKey:PutawayTaskID;constraint:OnDelete:CASCADE;" json:"putaway_task,omitempty"`
	GoodsReceiptLine    *GoodsReceiptLine `gorm:"foreignKey:GoodsReceiptLineID;constraint:OnDelete:RESTRICT;" json:"goods_receipt_line,omitempty"`
	Product             *Product          `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT;" json:"product,omitempty"`
	SourceLocation      *Location         `gorm:"foreignKey:SourceLocationID;constraint:OnDelete:RESTRICT;" json:"source_location,omitempty"`
	DestinationLocation *Location         `gorm:"foreignKey:DestinationLocationID;constraint:OnDelete:RESTRICT;" json:"destination_location,omitempty"`
	Batch               *ProductBatch     `gorm:"foreignKey:BatchID;constraint:OnDelete:SET NULL;" json:"batch,omitempty"`
}

func (PutawayTaskLine) TableName() string {
	return "putaway_task_lines"
}

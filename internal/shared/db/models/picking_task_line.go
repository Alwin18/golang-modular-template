package models

type PickingTaskLine struct {
	ID               uint    `gorm:"primary_key" json:"id"`
	PickingTaskID    uint    `gorm:"not null;index:idx_ptl_task" json:"picking_task_id"`
	SalesOrderLineID uint    `gorm:"not null;index:idx_ptl_sol" json:"sales_order_line_id"`
	ProductID        uint    `gorm:"not null;index:idx_ptl_product" json:"product_id"`
	FromLocationID   uint    `gorm:"not null;index:idx_ptl_fromloc" json:"from_location_id"`
	BatchID          *uint   `gorm:"index:idx_ptl_batch" json:"batch_id"`
	UOM              string  `gorm:"type:varchar(20);not null" json:"uom"`
	PlannedQty       float64 `gorm:"type:decimal(18,3);not null;column:planned_qty" json:"planned_qty"`
	PickedQty        float64 `gorm:"type:decimal(18,3);not null;default:0;column:picked_qty" json:"picked_qty"`
	SequenceNo       int     `gorm:"not null;column:sequence_no" json:"sequence_no"`

	// Relationships
	PickingTask    *PickingTask    `gorm:"foreignKey:PickingTaskID;constraint:OnDelete:CASCADE;" json:"picking_task,omitempty"`
	SalesOrderLine *SalesOrderLine `gorm:"foreignKey:SalesOrderLineID;constraint:OnDelete:RESTRICT;" json:"sales_order_line,omitempty"`
	Product        *Product        `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT;" json:"product,omitempty"`
	FromLocation   *Location       `gorm:"foreignKey:FromLocationID;constraint:OnDelete:RESTRICT;" json:"from_location,omitempty"`
	Batch          *ProductBatch   `gorm:"foreignKey:BatchID;constraint:OnDelete:SET NULL;" json:"batch,omitempty"`
}

func (PickingTaskLine) TableName() string {
	return "picking_task_lines"
}

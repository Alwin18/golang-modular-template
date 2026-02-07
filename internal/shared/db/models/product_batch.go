package models

import "time"

type ProductBatch struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	ProductID   uint       `gorm:"not null;index:idx_batch_product" json:"product_id"`
	BatchNumber string     `gorm:"type:varchar(100);not null;column:batch_number" json:"batch_number"`
	ExpiryDate  *time.Time `gorm:"type:date;column:expiry_date" json:"expiry_date"`

	// Relationships
	Product              *Product              `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;" json:"product,omitempty"`
	GoodsReceiptLines    []GoodsReceiptLine    `gorm:"foreignKey:BatchID" json:"goods_receipt_lines,omitempty"`
	InventoryBalances    []InventoryBalance    `gorm:"foreignKey:BatchID" json:"inventory_balances,omitempty"`
	InventoryMovements   []InventoryMovement   `gorm:"foreignKey:BatchID" json:"inventory_movements,omitempty"`
	PickingTaskLines     []PickingTaskLine     `gorm:"foreignKey:BatchID" json:"picking_task_lines,omitempty"`
	PutawayTaskLines     []PutawayTaskLine     `gorm:"foreignKey:BatchID" json:"putaway_task_lines,omitempty"`
	StockCountLines      []StockCountLine      `gorm:"foreignKey:BatchID" json:"stock_count_lines,omitempty"`
	StockAdjustmentLines []StockAdjustmentLine `gorm:"foreignKey:BatchID" json:"stock_adjustment_lines,omitempty"`
	SupplierReturnLines  []SupplierReturnLine  `gorm:"foreignKey:BatchID" json:"supplier_return_lines,omitempty"`
}

func (ProductBatch) TableName() string {
	return "product_batches"
}

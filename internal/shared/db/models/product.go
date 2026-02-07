package models

import "time"

type Product struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	SKU            string    `gorm:"type:varchar(100);not null;uniqueIndex:ux_products_sku" json:"sku"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"`
	Barcode        *string   `gorm:"type:varchar(100)" json:"barcode"`
	Description    *string   `gorm:"type:text" json:"description"`
	BaseUOM        string    `gorm:"type:varchar(20);not null;column:base_uom" json:"base_uom"`
	Weight         *float64  `gorm:"type:decimal(18,3)" json:"weight"`
	Volume         *float64  `gorm:"type:decimal(18,3)" json:"volume"`
	IsBatchManaged *bool     `gorm:"type:boolean;not null;default:false;column:is_batch_managed" json:"is_batch_managed"`
	IsSerialized   *bool     `gorm:"type:boolean;not null;default:false;column:is_serialized" json:"is_serialized"`
	IsActive       *bool     `gorm:"type:boolean;not null;default:true;column:is_active" json:"is_active"`
	CreatedAt      time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	ProductUOMs          []ProductUOM          `gorm:"foreignKey:ProductID" json:"product_uoms,omitempty"`
	ProductBatches       []ProductBatch        `gorm:"foreignKey:ProductID" json:"product_batches,omitempty"`
	PurchaseOrderLines   []PurchaseOrderLine   `gorm:"foreignKey:ProductID" json:"purchase_order_lines,omitempty"`
	GoodsReceiptLines    []GoodsReceiptLine    `gorm:"foreignKey:ProductID" json:"goods_receipt_lines,omitempty"`
	SalesOrderLines      []SalesOrderLine      `gorm:"foreignKey:ProductID" json:"sales_order_lines,omitempty"`
	PickingTaskLines     []PickingTaskLine     `gorm:"foreignKey:ProductID" json:"picking_task_lines,omitempty"`
	PutawayTaskLines     []PutawayTaskLine     `gorm:"foreignKey:ProductID" json:"putaway_task_lines,omitempty"`
	InventoryBalances    []InventoryBalance    `gorm:"foreignKey:ProductID" json:"inventory_balances,omitempty"`
	InventoryMovements   []InventoryMovement   `gorm:"foreignKey:ProductID" json:"inventory_movements,omitempty"`
	ShipmentPackageItems []ShipmentPackageItem `gorm:"foreignKey:ProductID" json:"shipment_package_items,omitempty"`
	StockCountLines      []StockCountLine      `gorm:"foreignKey:ProductID" json:"stock_count_lines,omitempty"`
	StockAdjustmentLines []StockAdjustmentLine `gorm:"foreignKey:ProductID" json:"stock_adjustment_lines,omitempty"`
	CustomerReturnLines  []CustomerReturnLine  `gorm:"foreignKey:ProductID" json:"customer_return_lines,omitempty"`
	SupplierReturnLines  []SupplierReturnLine  `gorm:"foreignKey:ProductID" json:"supplier_return_lines,omitempty"`
}

func (Product) TableName() string {
	return "products"
}

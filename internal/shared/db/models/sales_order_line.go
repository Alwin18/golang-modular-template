package models

type SalesOrderLine struct {
	ID           uint     `gorm:"primary_key" json:"id"`
	SalesOrderID uint     `gorm:"not null;index:idx_sol_so" json:"sales_order_id"`
	LineNo       int      `gorm:"not null;column:line_no" json:"line_no"`
	ProductID    uint     `gorm:"not null;index:idx_sol_product" json:"product_id"`
	UOM          string   `gorm:"type:varchar(20);not null" json:"uom"`
	OrderedQty   float64  `gorm:"type:decimal(18,3);not null;column:ordered_qty" json:"ordered_qty"`
	AllocatedQty float64  `gorm:"type:decimal(18,3);not null;default:0;column:allocated_qty" json:"allocated_qty"`
	ShippedQty   float64  `gorm:"type:decimal(18,3);not null;default:0;column:shipped_qty" json:"shipped_qty"`
	UnitPrice    *float64 `gorm:"type:decimal(18,2);column:unit_price" json:"unit_price"`

	// Relationships
	SalesOrder           *SalesOrder           `gorm:"foreignKey:SalesOrderID;constraint:OnDelete:CASCADE;" json:"sales_order,omitempty"`
	Product              *Product              `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT;" json:"product,omitempty"`
	PickingTaskLines     []PickingTaskLine     `gorm:"foreignKey:SalesOrderLineID" json:"picking_task_lines,omitempty"`
	ShipmentPackageItems []ShipmentPackageItem `gorm:"foreignKey:SalesOrderLineID" json:"shipment_package_items,omitempty"`
	CustomerReturnLines  []CustomerReturnLine  `gorm:"foreignKey:SalesOrderLineID" json:"customer_return_lines,omitempty"`
}

func (SalesOrderLine) TableName() string {
	return "sales_order_lines"
}

package models

type CustomerReturnLine struct {
	ID               uint    `gorm:"primary_key" json:"id"`
	CustomerReturnID uint    `gorm:"not null;index:idx_crlines_cr" json:"customer_return_id"`
	SalesOrderLineID uint    `gorm:"not null;index:idx_crlines_sol" json:"sales_order_line_id"`
	ProductID        uint    `gorm:"not null;index:idx_crlines_product" json:"product_id"`
	UOM              string  `gorm:"type:varchar(20);not null" json:"uom"`
	ReturnedQty      float64 `gorm:"type:decimal(18,3);not null;column:returned_qty" json:"returned_qty"`
	QCStatus         *string `gorm:"type:varchar(50);column:qc_status" json:"qc_status"`
	ReturnReasonCode *string `gorm:"type:varchar(50);column:return_reason_code" json:"return_reason_code"`

	// Relationships
	CustomerReturn *CustomerReturn `gorm:"foreignKey:CustomerReturnID;constraint:OnDelete:CASCADE;" json:"customer_return,omitempty"`
	SalesOrderLine *SalesOrderLine `gorm:"foreignKey:SalesOrderLineID;constraint:OnDelete:RESTRICT;" json:"sales_order_line,omitempty"`
	Product        *Product        `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT;" json:"product,omitempty"`
}

func (CustomerReturnLine) TableName() string {
	return "customer_return_lines"
}

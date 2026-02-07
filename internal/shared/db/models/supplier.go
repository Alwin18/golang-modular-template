package models

import "time"

type Supplier struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Code      string    `gorm:"type:varchar(50);not null;uniqueIndex:ux_suppliers_code" json:"code"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Address   *string   `gorm:"type:text" json:"address"`
	City      *string   `gorm:"type:varchar(100)" json:"city"`
	Country   *string   `gorm:"type:varchar(100)" json:"country"`
	Phone     *string   `gorm:"type:varchar(50)" json:"phone"`
	Email     *string   `gorm:"type:varchar(255)" json:"email"`
	TaxID     *string   `gorm:"type:varchar(100);column:tax_id" json:"tax_id"`
	IsActive  *bool     `gorm:"type:boolean;not null;default:true;column:is_active" json:"is_active"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	PurchaseOrders  []PurchaseOrder  `gorm:"foreignKey:SupplierID" json:"purchase_orders,omitempty"`
	GoodsReceipts   []GoodsReceipt   `gorm:"foreignKey:SupplierID" json:"goods_receipts,omitempty"`
	SupplierReturns []SupplierReturn `gorm:"foreignKey:SupplierID" json:"supplier_returns,omitempty"`
}

func (Supplier) TableName() string {
	return "suppliers"
}

package models

import "time"

type Customer struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Code      string    `gorm:"type:varchar(50);not null;uniqueIndex:ux_customers_code" json:"code"`
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
	SalesOrders     []SalesOrder     `gorm:"foreignKey:CustomerID" json:"sales_orders,omitempty"`
	CustomerReturns []CustomerReturn `gorm:"foreignKey:CustomerID" json:"customer_returns,omitempty"`
}

func (Customer) TableName() string {
	return "customers"
}

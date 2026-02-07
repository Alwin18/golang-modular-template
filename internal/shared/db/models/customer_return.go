package models

import "time"

type CustomerReturn struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	ReturnNumber string     `gorm:"type:varchar(100);not null;uniqueIndex:ux_custret_number;column:return_number" json:"return_number"`
	SalesOrderID uint       `gorm:"not null;index:idx_custret_so" json:"sales_order_id"`
	CustomerID   uint       `gorm:"not null;index:idx_custret_customer" json:"customer_id"`
	WarehouseID  uint       `gorm:"not null;index:idx_custret_wh" json:"warehouse_id"`
	Status       string     `gorm:"type:varchar(50);not null" json:"status"`
	Reason       *string    `gorm:"type:text" json:"reason"`
	RequestedAt  time.Time  `gorm:"not null;column:requested_at" json:"requested_at"`
	ReceivedAt   *time.Time `gorm:"column:received_at" json:"received_at"`

	// Relationships
	SalesOrder *SalesOrder          `gorm:"foreignKey:SalesOrderID;constraint:OnDelete:RESTRICT;" json:"sales_order,omitempty"`
	Customer   *Customer            `gorm:"foreignKey:CustomerID;constraint:OnDelete:RESTRICT;" json:"customer,omitempty"`
	Warehouse  *Warehouse           `gorm:"foreignKey:WarehouseID;constraint:OnDelete:RESTRICT;" json:"warehouse,omitempty"`
	Lines      []CustomerReturnLine `gorm:"foreignKey:CustomerReturnID" json:"lines,omitempty"`
}

func (CustomerReturn) TableName() string {
	return "customer_returns"
}

package models

import "time"

type SalesOrder struct {
	ID                uint       `gorm:"primary_key" json:"id"`
	SONumber          string     `gorm:"type:varchar(100);not null;uniqueIndex:ux_so_number;column:so_number" json:"so_number"`
	ExternalRef       *string    `gorm:"type:varchar(255);column:external_ref" json:"external_ref"`
	CustomerID        uint       `gorm:"not null;index:idx_so_customer" json:"customer_id"`
	WarehouseID       uint       `gorm:"not null;index:idx_so_warehouse" json:"warehouse_id"`
	Status            string     `gorm:"type:varchar(50);not null" json:"status"`
	OrderDate         time.Time  `gorm:"not null;column:order_date" json:"order_date"`
	RequestedShipDate *time.Time `gorm:"column:requested_ship_date" json:"requested_ship_date"`
	Priority          *string    `gorm:"type:varchar(20)" json:"priority"`
	ShippingAddress   *string    `gorm:"type:text;column:shipping_address" json:"shipping_address"`
	ShippingCity      *string    `gorm:"type:varchar(100);column:shipping_city" json:"shipping_city"`
	ShippingCountry   *string    `gorm:"type:varchar(100);column:shipping_country" json:"shipping_country"`
	ShippingPhone     *string    `gorm:"type:varchar(50);column:shipping_phone" json:"shipping_phone"`
	CreatedAt         time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	Customer        *Customer        `gorm:"foreignKey:CustomerID;constraint:OnDelete:RESTRICT;" json:"customer,omitempty"`
	Warehouse       *Warehouse       `gorm:"foreignKey:WarehouseID;constraint:OnDelete:RESTRICT;" json:"warehouse,omitempty"`
	Lines           []SalesOrderLine `gorm:"foreignKey:SalesOrderID" json:"lines,omitempty"`
	PickingTasks    []PickingTask    `gorm:"foreignKey:SalesOrderID" json:"picking_tasks,omitempty"`
	ShipmentOrders  []ShipmentOrder  `gorm:"foreignKey:SalesOrderID" json:"shipment_orders,omitempty"`
	CustomerReturns []CustomerReturn `gorm:"foreignKey:SalesOrderID" json:"customer_returns,omitempty"`
}

func (SalesOrder) TableName() string {
	return "sales_orders"
}

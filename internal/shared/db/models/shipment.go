package models

import "time"

type Shipment struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	ShipmentNumber string     `gorm:"type:varchar(100);not null;uniqueIndex:ux_shipment_number;column:shipment_number" json:"shipment_number"`
	WarehouseID    uint       `gorm:"not null;index:idx_ship_wh" json:"warehouse_id"`
	CarrierID      *uint      `gorm:"index:idx_ship_carrier" json:"carrier_id"`
	Status         string     `gorm:"type:varchar(50);not null" json:"status"`
	DispatchTime   *time.Time `gorm:"column:dispatch_time" json:"dispatch_time"`
	DeliveredTime  *time.Time `gorm:"column:delivered_time" json:"delivered_time"`
	CreatedAt      time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`

	// Relationships
	Warehouse        *Warehouse        `gorm:"foreignKey:WarehouseID;constraint:OnDelete:RESTRICT;" json:"warehouse,omitempty"`
	Carrier          *Carrier          `gorm:"foreignKey:CarrierID;constraint:OnDelete:SET NULL;" json:"carrier,omitempty"`
	ShipmentOrders   []ShipmentOrder   `gorm:"foreignKey:ShipmentID" json:"shipment_orders,omitempty"`
	ShipmentPackages []ShipmentPackage `gorm:"foreignKey:ShipmentID" json:"shipment_packages,omitempty"`
}

func (Shipment) TableName() string {
	return "shipments"
}

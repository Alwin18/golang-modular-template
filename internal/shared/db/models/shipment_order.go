package models

type ShipmentOrder struct {
	ID           uint `gorm:"primary_key" json:"id"`
	ShipmentID   uint `gorm:"not null;index:idx_shiporder_ship" json:"shipment_id"`
	SalesOrderID uint `gorm:"not null;index:idx_shiporder_so" json:"sales_order_id"`

	// Relationships
	Shipment   *Shipment   `gorm:"foreignKey:ShipmentID;constraint:OnDelete:CASCADE;" json:"shipment,omitempty"`
	SalesOrder *SalesOrder `gorm:"foreignKey:SalesOrderID;constraint:OnDelete:RESTRICT;" json:"sales_order,omitempty"`
}

func (ShipmentOrder) TableName() string {
	return "shipment_orders"
}

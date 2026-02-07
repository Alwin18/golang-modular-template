package models

type ShipmentPackageItem struct {
	ID                uint    `gorm:"primary_key" json:"id"`
	ShipmentPackageID uint    `gorm:"not null;index:idx_pkgitems_pkg" json:"shipment_package_id"`
	SalesOrderLineID  uint    `gorm:"not null;index:idx_pkgitems_sol" json:"sales_order_line_id"`
	ProductID         uint    `gorm:"not null;index:idx_pkgitems_product" json:"product_id"`
	UOM               string  `gorm:"type:varchar(20);not null" json:"uom"`
	Qty               float64 `gorm:"type:decimal(18,3);not null" json:"qty"`

	// Relationships
	ShipmentPackage *ShipmentPackage `gorm:"foreignKey:ShipmentPackageID;constraint:OnDelete:CASCADE;" json:"shipment_package,omitempty"`
	SalesOrderLine  *SalesOrderLine  `gorm:"foreignKey:SalesOrderLineID;constraint:OnDelete:RESTRICT;" json:"sales_order_line,omitempty"`
	Product         *Product         `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT;" json:"product,omitempty"`
}

func (ShipmentPackageItem) TableName() string {
	return "shipment_package_items"
}

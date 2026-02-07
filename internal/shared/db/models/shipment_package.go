package models

type ShipmentPackage struct {
	ID             uint     `gorm:"primary_key" json:"id"`
	ShipmentID     uint     `gorm:"not null;index:idx_pkg_ship" json:"shipment_id"`
	PackageNumber  string   `gorm:"type:varchar(100);not null;uniqueIndex:ux_package_number;column:package_number" json:"package_number"`
	TrackingNumber *string  `gorm:"type:varchar(255);column:tracking_number" json:"tracking_number"`
	Weight         *float64 `gorm:"type:decimal(18,3)" json:"weight"`
	Volume         *float64 `gorm:"type:decimal(18,3)" json:"volume"`

	// Relationships
	Shipment             *Shipment             `gorm:"foreignKey:ShipmentID;constraint:OnDelete:CASCADE;" json:"shipment,omitempty"`
	ShipmentPackageItems []ShipmentPackageItem `gorm:"foreignKey:ShipmentPackageID" json:"shipment_package_items,omitempty"`
}

func (ShipmentPackage) TableName() string {
	return "shipment_packages"
}

package models

type WarehouseZone struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	WarehouseID uint   `gorm:"not null;index:idx_zones_wh" json:"warehouse_id"`
	Code        string `gorm:"type:varchar(50);not null" json:"code"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	ZoneType    string `gorm:"type:varchar(50);not null;column:zone_type" json:"zone_type"`
	IsActive    *bool  `gorm:"type:boolean;not null;default:true;column:is_active" json:"is_active"`

	// Relationships
	Warehouse Warehouse  `gorm:"foreignKey:WarehouseID;constraint:OnDelete:CASCADE;" json:"warehouse,omitempty"`
	Locations []Location `gorm:"foreignKey:ZoneID" json:"locations,omitempty"`
}

func (WarehouseZone) TableName() string {
	return "warehouse_zones"
}

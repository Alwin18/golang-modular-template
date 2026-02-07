package models

type Carrier struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Code     string `gorm:"type:varchar(50);not null;uniqueIndex:ux_carriers_code" json:"code"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	IsActive *bool  `gorm:"type:boolean;not null;default:true;column:is_active" json:"is_active"`

	// Relationships
	Shipments []Shipment `gorm:"foreignKey:CarrierID" json:"shipments,omitempty"`
}

func (Carrier) TableName() string {
	return "carriers"
}

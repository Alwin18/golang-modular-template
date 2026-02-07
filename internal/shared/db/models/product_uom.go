package models

type ProductUOM struct {
	ID                uint    `gorm:"primary_key" json:"id"`
	ProductID         uint    `gorm:"not null;index:idx_prod_uom" json:"product_id"`
	UOM               string  `gorm:"type:varchar(20);not null" json:"uom"`
	ConversionToBase  float64 `gorm:"type:decimal(18,6);not null;column:conversion_to_base" json:"conversion_to_base"`
	IsDefaultSales    *bool   `gorm:"type:boolean;not null;default:false;column:is_default_sales" json:"is_default_sales"`
	IsDefaultPurchase *bool   `gorm:"type:boolean;not null;default:false;column:is_default_purchase" json:"is_default_purchase"`

	// Relationships
	Product *Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;" json:"product,omitempty"`
}

func (ProductUOM) TableName() string {
	return "product_uoms"
}

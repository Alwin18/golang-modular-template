package models

type Location struct {
	ID           uint     `gorm:"primary_key" json:"id"`
	WarehouseID  uint     `gorm:"not null;index:idx_locations_wh" json:"warehouse_id"`
	ZoneID       uint     `gorm:"not null;index:idx_locations_zone" json:"zone_id"`
	Code         string   `gorm:"type:varchar(50);not null" json:"code"`
	Description  *string  `gorm:"type:text" json:"description"`
	LocationType string   `gorm:"type:varchar(50);not null;column:location_type" json:"location_type"`
	MaxVolume    *float64 `gorm:"type:decimal(18,2);column:max_volume" json:"max_volume"`
	MaxWeight    *float64 `gorm:"type:decimal(18,2);column:max_weight" json:"max_weight"`
	IsActive     *bool    `gorm:"type:boolean;not null;default:true;column:is_active" json:"is_active"`

	// Relationships
	Warehouse            *Warehouse            `gorm:"foreignKey:WarehouseID;constraint:OnDelete:CASCADE;" json:"warehouse,omitempty"`
	Zone                 *WarehouseZone        `gorm:"foreignKey:ZoneID;constraint:OnDelete:CASCADE;" json:"zone,omitempty"`
	InventoryBalances    []InventoryBalance    `gorm:"foreignKey:LocationID" json:"inventory_balances,omitempty"`
	MovementsFrom        []InventoryMovement   `gorm:"foreignKey:FromLocationID" json:"movements_from,omitempty"`
	MovementsTo          []InventoryMovement   `gorm:"foreignKey:ToLocationID" json:"movements_to,omitempty"`
	PickingTaskLines     []PickingTaskLine     `gorm:"foreignKey:FromLocationID" json:"picking_task_lines,omitempty"`
	PutawayLinesSource   []PutawayTaskLine     `gorm:"foreignKey:SourceLocationID" json:"putaway_lines_source,omitempty"`
	PutawayLinesDest     []PutawayTaskLine     `gorm:"foreignKey:DestinationLocationID" json:"putaway_lines_dest,omitempty"`
	StockCountLines      []StockCountLine      `gorm:"foreignKey:LocationID" json:"stock_count_lines,omitempty"`
	StockAdjustmentLines []StockAdjustmentLine `gorm:"foreignKey:LocationID" json:"stock_adjustment_lines,omitempty"`
}

func (Location) TableName() string {
	return "locations"
}

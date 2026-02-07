package models

import "time"

type Warehouse struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Code      string    `gorm:"type:varchar(50);not null;uniqueIndex:ux_warehouses_code" json:"code"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Address   *string   `gorm:"type:text" json:"address"`
	City      *string   `gorm:"type:varchar(100)" json:"city"`
	Country   *string   `gorm:"type:varchar(100)" json:"country"`
	TimeZone  *string   `gorm:"type:varchar(50);column:time_zone" json:"time_zone"`
	IsActive  *bool     `gorm:"type:boolean;not null;default:true;column:is_active" json:"is_active"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	Zones              []WarehouseZone     `gorm:"foreignKey:WarehouseID" json:"zones,omitempty"`
	Locations          []Location          `gorm:"foreignKey:WarehouseID" json:"locations,omitempty"`
	PurchaseOrders     []PurchaseOrder     `gorm:"foreignKey:WarehouseID" json:"purchase_orders,omitempty"`
	GoodsReceipts      []GoodsReceipt      `gorm:"foreignKey:WarehouseID" json:"goods_receipts,omitempty"`
	SalesOrders        []SalesOrder        `gorm:"foreignKey:WarehouseID" json:"sales_orders,omitempty"`
	PickingWaves       []PickingWave       `gorm:"foreignKey:WarehouseID" json:"picking_waves,omitempty"`
	PickingTasks       []PickingTask       `gorm:"foreignKey:WarehouseID" json:"picking_tasks,omitempty"`
	PutawayTasks       []PutawayTask       `gorm:"foreignKey:WarehouseID" json:"putaway_tasks,omitempty"`
	InventoryBalances  []InventoryBalance  `gorm:"foreignKey:WarehouseID" json:"inventory_balances,omitempty"`
	InventoryMovements []InventoryMovement `gorm:"foreignKey:WarehouseID" json:"inventory_movements,omitempty"`
	Shipments          []Shipment          `gorm:"foreignKey:WarehouseID" json:"shipments,omitempty"`
	StockCounts        []StockCount        `gorm:"foreignKey:WarehouseID" json:"stock_counts,omitempty"`
	StockAdjustments   []StockAdjustment   `gorm:"foreignKey:WarehouseID" json:"stock_adjustments,omitempty"`
	CustomerReturns    []CustomerReturn    `gorm:"foreignKey:WarehouseID" json:"customer_returns,omitempty"`
	SupplierReturns    []SupplierReturn    `gorm:"foreignKey:WarehouseID" json:"supplier_returns,omitempty"`
}

func (Warehouse) TableName() string {
	return "warehouses"
}

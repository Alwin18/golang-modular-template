package models

import "time"

type PickingTask struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	PickingWaveID *uint      `gorm:"index:idx_pt_wave" json:"picking_wave_id"`
	SalesOrderID  uint       `gorm:"not null;index:idx_pt_so" json:"sales_order_id"`
	WarehouseID   uint       `gorm:"not null;index:idx_pt_wh" json:"warehouse_id"`
	AssignedTo    *uint      `gorm:"index:idx_pt_assigned_to" json:"assigned_to"`
	Status        string     `gorm:"type:varchar(50);not null" json:"status"`
	CreatedAt     time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	StartedAt     *time.Time `gorm:"column:started_at" json:"started_at"`
	CompletedAt   *time.Time `gorm:"column:completed_at" json:"completed_at"`

	// Relationships
	PickingWave *PickingWave      `gorm:"foreignKey:PickingWaveID;constraint:OnDelete:SET NULL;" json:"picking_wave,omitempty"`
	SalesOrder  *SalesOrder       `gorm:"foreignKey:SalesOrderID;constraint:OnDelete:RESTRICT;" json:"sales_order,omitempty"`
	Warehouse   *Warehouse        `gorm:"foreignKey:WarehouseID;constraint:OnDelete:RESTRICT;" json:"warehouse,omitempty"`
	Assignee    *User             `gorm:"foreignKey:AssignedTo;constraint:OnDelete:SET NULL;" json:"assignee,omitempty"`
	Lines       []PickingTaskLine `gorm:"foreignKey:PickingTaskID" json:"lines,omitempty"`
}

func (PickingTask) TableName() string {
	return "picking_tasks"
}

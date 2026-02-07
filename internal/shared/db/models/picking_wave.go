package models

import "time"

type PickingWave struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	WaveNumber  string    `gorm:"type:varchar(100);not null;uniqueIndex:ux_wave_number;column:wave_number" json:"wave_number"`
	WarehouseID uint      `gorm:"not null;index:idx_wave_wh" json:"warehouse_id"`
	Status      string    `gorm:"type:varchar(50);not null" json:"status"`
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy   uint      `gorm:"not null;index:idx_wave_created_by" json:"created_by"`

	// Relationships
	Warehouse    *Warehouse    `gorm:"foreignKey:WarehouseID;constraint:OnDelete:RESTRICT;" json:"warehouse,omitempty"`
	Creator      *User         `gorm:"foreignKey:CreatedBy;constraint:OnDelete:RESTRICT;" json:"creator,omitempty"`
	PickingTasks []PickingTask `gorm:"foreignKey:PickingWaveID" json:"picking_tasks,omitempty"`
}

func (PickingWave) TableName() string {
	return "picking_waves"
}

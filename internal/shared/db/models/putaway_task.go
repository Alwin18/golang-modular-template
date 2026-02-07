package models

import "time"

type PutawayTask struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	WarehouseID    uint       `gorm:"not null;index:idx_putaway_wh" json:"warehouse_id"`
	GoodsReceiptID uint       `gorm:"not null;index:idx_putaway_gr" json:"goods_receipt_id"`
	AssignedTo     *uint      `gorm:"index:idx_putaway_assigned_to" json:"assigned_to"`
	Status         string     `gorm:"type:varchar(50);not null" json:"status"`
	CreatedAt      time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	StartedAt      *time.Time `gorm:"column:started_at" json:"started_at"`
	CompletedAt    *time.Time `gorm:"column:completed_at" json:"completed_at"`

	// Relationships
	Warehouse    *Warehouse        `gorm:"foreignKey:WarehouseID;constraint:OnDelete:RESTRICT;" json:"warehouse,omitempty"`
	GoodsReceipt *GoodsReceipt     `gorm:"foreignKey:GoodsReceiptID;constraint:OnDelete:RESTRICT;" json:"goods_receipt,omitempty"`
	Assignee     *User             `gorm:"foreignKey:AssignedTo;constraint:OnDelete:SET NULL;" json:"assignee,omitempty"`
	Lines        []PutawayTaskLine `gorm:"foreignKey:PutawayTaskID" json:"lines,omitempty"`
}

func (PutawayTask) TableName() string {
	return "putaway_tasks"
}

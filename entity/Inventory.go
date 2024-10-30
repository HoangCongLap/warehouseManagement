package entity

import (
	"time"
)

type Inventory struct {
	InventoryID int       `json:"inventory_id" gorm:"primaryKey;autoIncrement"`
	ProductID   int       `json:"product_id" gorm:"not null"`
	WarehouseID int       `json:"warehouse_id" gorm:"not null"`
	Quantity    int       `json:"quantity" gorm:"not null"`
	LastUpdated time.Time `json:"last_updated" gorm:"autoUpdateTime"`
}

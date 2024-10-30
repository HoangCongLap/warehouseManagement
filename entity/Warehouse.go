package entity

type Warehouse struct {
	WarehouseID int    `json:"warehouse_id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"size:255;not null"`
	Location    string `json:"location" gorm:"size:255"`
	Capacity    int    `json:"capacity"`
}

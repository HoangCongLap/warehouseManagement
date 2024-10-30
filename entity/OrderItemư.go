package entity

type OrderItem struct {
	OrderItemID int     `json:"order_item_id" gorm:"primaryKey;autoIncrement"`
	OrderID     int     `json:"order_id" gorm:"not null"`
	ProductID   int     `json:"product_id" gorm:"not null"`
	Quantity    int     `json:"quantity" gorm:"not null"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null"`
}

package entity

import (
	"time"
)

type Order struct {
	OrderID      int       `json:"order_id" gorm:"primaryKey;autoIncrement"`
	CustomerName string    `json:"customer_name" gorm:"size:255;not null"`
	OrderDate    time.Time `json:"order_date" gorm:"autoCreateTime"`
	Status       string    `json:"status" gorm:"type:enum('pending', 'shipped', 'completed', 'canceled');default:'pending'"`
}

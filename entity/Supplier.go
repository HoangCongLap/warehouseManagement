package entity

type Supplier struct {
	SupplierID  int    `json:"supplier_id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"size:255;not null"`
	ContactInfo string `json:"contact_info" gorm:"size:255"`
}

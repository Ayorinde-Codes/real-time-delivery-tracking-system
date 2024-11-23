package models

// Order model for order data
type Order struct {
	ID         uint `gorm:"primaryKey"`
	CustomerID uint
	Status     string
}

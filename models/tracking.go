package models

// Tracking model for order tracking data
type Tracking struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	Latitude  float64
	Longitude float64
}

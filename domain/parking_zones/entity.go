package parkingzones

import (
	"time"
)

type ParkingZone struct {
	ID             uint      `gorm:"primaryKey"`
	Name           string    `gorm:"not null"`
	Type           string    `gorm:"type:varchar(20);not null"`
	Total_Capacity int       `gorm:"not null"`
	Price_per_Hour float64   `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

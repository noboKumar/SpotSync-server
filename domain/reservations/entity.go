package reservations

import (
	"time"

	parkingzones "github.com/noboKumar/SpotSync-server/domain/parking_zones"
	"github.com/noboKumar/SpotSync-server/domain/users"
)

type Reservation struct {
	ID uint `gorm:"primaryKey"`

	UserID uint       `gorm:"not null"`
	User   users.User `gorm:"foreignKey:UserID"`

	ZoneID uint                     `gorm:"not null"`
	Zone   parkingzones.ParkingZone `gorm:"foreignKey:ZoneID"`

	LicensePlate string `gorm:"type:varchar(15) not null"`
	Status      string `gorm:"type:varchar(20);default:'active';check:status IN ('active','completed','cancelled')"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

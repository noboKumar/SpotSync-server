package parkingzones

import "gorm.io/gorm"

type Repository interface {
	CreateParkingZone(parkingZone *ParkingZone) error
}

type repository struct {
	db *gorm.DB
}

func NewParkingZoneRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateParkingZone(parkingZone *ParkingZone) error {
	return r.db.Create(parkingZone).Error
}


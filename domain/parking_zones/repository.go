package parkingzones

import "gorm.io/gorm"

type Repository interface {
	CreateParkingZone(parkingZone *ParkingZone) error
	GetAllParkingZone(parkingZones *[]ParkingZone) error
	GetSingleParkingZone(id string) (*ParkingZone, error)
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

func (r *repository) GetAllParkingZone(parkingZone *[]ParkingZone) error {
	return r.db.Find(parkingZone).Error
}

func (r *repository) GetSingleParkingZone(id string) (*ParkingZone, error) {
	var parkingZone ParkingZone
	err := r.db.First(&parkingZone, id).Error
	if err != nil {
		return nil, err
	}
	return &parkingZone, nil
}

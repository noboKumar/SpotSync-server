package reservations

import "gorm.io/gorm"

type Repository interface {
	// reserve parking zone

}

type repository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) ReserveParkingZone(reservation *Reservation) error {
	result := r.db.Create(reservation)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

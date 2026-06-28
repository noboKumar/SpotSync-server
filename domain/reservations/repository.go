package reservations

import "gorm.io/gorm"

type Repository interface {
	ReserveParkingZone(
		reservation *Reservation,
	) error
	GetMyReservation(userID uint) ([]Reservation, error)
}

type repository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) ReserveParkingZone(
	reservation *Reservation,
) error {

	result := r.db.Create(reservation)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repository) GetMyReservation(userID uint) ([]Reservation, error) {
	var reservations []Reservation

	result := r.db.
		Preload("Zone").
		Where("user_id = ?", userID).
		Find(&reservations)

	if result.Error != nil {
		return nil, result.Error
	}

	return reservations, nil
}

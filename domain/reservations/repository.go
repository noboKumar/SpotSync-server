package reservations

import (
	"gorm.io/gorm"
)

type Repository interface {
	ReserveParkingZone(
		reservation *Reservation,
	) error
	GetMyReservation(userID uint) ([]Reservation, error)
	CancelReservation(reservationID uint, userID uint) error
	GetAllReservations(reservations *[]Reservation) error
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

func (r *repository) CancelReservation(reservationID uint, userID uint) error {
	result := r.db.
		Model(&Reservation{}).
		Where("id = ? AND user_id = ?", reservationID, userID).
		Update("status", "cancelled")

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *repository) GetAllReservations(reservations *[]Reservation) error {

	result := r.db.Preload("Zone").Preload("User").Find(&reservations)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package reservations

import (
	"errors"

	"github.com/noboKumar/SpotSync-server/domain/reservations/dto"
)

var ErrZoneFull = errors.New("parking zone is full")

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) ReserveParkingZone(
	req dto.ReserveParkingZoneReq,
	userID uint,
) (dto.ReserveParkingZoneRes, error) {

	reservation := Reservation{
		ZoneID:       req.ZoneId,
		LicensePlate: req.LicensePlate,
		Status:       "active",
		UserID:       userID,
	}

	err := s.repo.ReserveParkingZone(&reservation)

	if errors.Is(err, ErrZoneFull) {
		return dto.ReserveParkingZoneRes{}, ErrZoneFull
	}

	if err != nil {
		return dto.ReserveParkingZoneRes{}, err
	}

	return dto.ReserveParkingZoneRes{
		ID:           reservation.ID,
		UserID:       reservation.UserID,
		ZoneID:       reservation.ZoneID,
		LicensePlate: reservation.LicensePlate,
		Status:       reservation.Status,
		CreatedAt:    reservation.CreatedAt,
		UpdatedAt:    reservation.UpdatedAt,
	}, nil
}
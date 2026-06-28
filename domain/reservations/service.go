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

func (s *service) GetMyReservation(userID uint) ([]dto.MyReservationResponse, error) {
	reservations, err := s.repo.GetMyReservation(userID)
	if err != nil {
		return nil, err
	}

	var res []dto.MyReservationResponse
	for _, r := range reservations {
		res = append(res, dto.MyReservationResponse{
			ID:           r.ID,
			LicensePlate: r.LicensePlate,
			Status:       r.Status,
			Zone: dto.ZoneResponse{
				ID:   r.Zone.ID,
				Name: r.Zone.Name,
				Type: r.Zone.Type,
			},
			CreatedAt: r.CreatedAt,
		})
	}

	return res, nil
}

func (s *service) CancelReservation(reservationID uint, userID uint) (dto.DeleteResponse, error) {
	err := s.repo.CancelReservation(reservationID, userID)
	if err != nil {
		return dto.DeleteResponse{Success: false, Message: "Failed to cancel reservation"}, err
	}
	return dto.DeleteResponse{Success: true, Message: "Reservation cancelled successfully"}, nil

}

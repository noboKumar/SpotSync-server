package reservations

import "github.com/noboKumar/SpotSync-server/domain/reservations/dto"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) ReserveParkingZone(req dto.ReserveParkingZoneReq) (dto.ReserveParkingZoneRes, error) {
	reservation := Reservation{
		ZoneID:       req.ZoneId,
		LicensePlate: req.LicensePlate,
		Status:       "active",
		UserID:       req.UserID,
	}

	err := s.repo.ReserveParkingZone(&reservation)

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

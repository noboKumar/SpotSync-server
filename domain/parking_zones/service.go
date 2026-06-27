package parkingzones

import "github.com/noboKumar/SpotSync-server/domain/parking_zones/dto"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateParkingZone(req dto.CreateParkingZoneRequest) (*dto.CreateParkingZoneResponse, error) {
	parkingZone := ParkingZone{
		Name:           req.Name,
		Type:           req.Type,
		Total_Capacity: req.Total_Capacity,
		Price_per_Hour: req.Price_per_Hour,
	}
	err := s.repo.CreateParkingZone(&parkingZone)

	if err != nil {
		return nil, err
	}

	return &dto.CreateParkingZoneResponse{
		ID:             parkingZone.ID,
		Name:           parkingZone.Name,
		Type:           parkingZone.Type,
		Total_Capacity: parkingZone.Total_Capacity,
		Price_per_Hour: parkingZone.Price_per_Hour,
		CreatedAt:      parkingZone.CreatedAt,
		UpdatedAt:      parkingZone.UpdatedAt,
	}, nil
}

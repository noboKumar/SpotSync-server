package dto

type CreateParkingZoneRequest struct {
	Name           string  `json:"name" validate:"required"`
	Type           string  `json:"type" validate:"required,oneof=public private"`
	Total_Capacity int     `json:"total_capacity" validate:"required,min=1"`
	Price_per_Hour float64 `json:"price_per_hour" validate:"required,gt=0"`
}

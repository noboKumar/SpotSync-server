package dto

type ReserveParkingZoneReq struct {
	ZoneId       uint   `json:"zone_id" validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required"`
	UserID       uint   `json:"-"`
}

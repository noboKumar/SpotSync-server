package parkingzones

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	parkingZoneRepository := NewParkingZoneRepository(db)
	parkingZoneService := NewService(parkingZoneRepository)
	parkingZoneHandler := NewHandler(parkingZoneService)

	// Set up the API group for parking zone-related routes
	api := e.Group("/api/v1")

	// Register the route for creating a parking zone
	api.POST("/zones", parkingZoneHandler.CreateParkingZone)
}

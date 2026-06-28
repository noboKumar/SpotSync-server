package parkingzones

import (
	"github.com/labstack/echo/v5"
	"github.com/noboKumar/SpotSync-server/auth"
	"github.com/noboKumar/SpotSync-server/config"
	"github.com/noboKumar/SpotSync-server/middleware"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	parkingZoneRepository := NewParkingZoneRepository(db)
	parkingZoneService := NewService(parkingZoneRepository)
	parkingZoneHandler := NewHandler(parkingZoneService)
	jwtService := auth.NewJwtService(cfg.JwtSecret)

	// Set up the API group for parking zone-related routes
	api := e.Group("/api/v1")

	// Register the route for creating a parking zone
	api.POST("/zones",
		parkingZoneHandler.CreateParkingZone,
		middleware.AuthMiddleware(jwtService),
		middleware.AdminMiddleware,
	)

	// Register the route for getting all parking zones
	api.GET("/zones", parkingZoneHandler.GetAllParkingZone)
}

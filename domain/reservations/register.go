package reservations

import (
	"github.com/labstack/echo/v5"
	"github.com/noboKumar/SpotSync-server/auth"
	"github.com/noboKumar/SpotSync-server/config"
	"github.com/noboKumar/SpotSync-server/middleware"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	reservationRepository := NewReservationRepository(db)

	reservationService := NewService(reservationRepository)

	reservationHandler := NewHandler(reservationService)

	jwtService := auth.NewJwtService(cfg.JwtSecret)

	api := e.Group("/api/v1")

	api.POST("/reservations", reservationHandler.ReserveParkingZone, middleware.AuthMiddleware(jwtService))
	api.GET("/reservations/my-reservations", reservationHandler.GetMyReservation, middleware.AuthMiddleware(jwtService))
}

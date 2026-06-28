package server

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"

	"github.com/noboKumar/SpotSync-server/config"
	parkingzones "github.com/noboKumar/SpotSync-server/domain/parking_zones"
	"github.com/noboKumar/SpotSync-server/domain/reservations"
	"github.com/noboKumar/SpotSync-server/domain/users"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Start(cfg *config.Config, db *gorm.DB) {
	db.AutoMigrate(&users.User{}, &parkingzones.ParkingZone{}, &reservations.Reservation{})

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	//middleware
	e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"message": "SpotSync API is running...",
		})
	})

	e.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	//routes
	users.RegisterRoutes(e, db, cfg)
	parkingzones.RegisterRoutes(e, db, cfg)
	reservations.RegisterRoutes(e, db, cfg)

	e.Start(":" + cfg.Port)
}

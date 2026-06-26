package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"

	"github.com/noboKumar/SpotSync-server/config"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Start(cfg *config.Config, db *gorm.DB) {
	db.AutoMigrate()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	//middleware
	e.GET("/health", func(c *echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "ok",
		})
	})

}

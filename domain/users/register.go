package users

import (
	"github.com/noboKumar/SpotSync-server/auth"
	"github.com/noboKumar/SpotSync-server/config"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

// RegisterRoutes sets up the routes for user-related operations, including registration.
func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	userRepository := NewUserRepository(db)               //  user repository
	jwtService := auth.NewJwtService(cfg.JwtSecret)       //  JWT service with the provided secret
	userService := NewService(userRepository, jwtService) // user service with the repository and JWT service
	userHandler := NewHandler(userService)            // user handler with the service

	// Set up the API group for user-related routes
	api := e.Group("/api/v1/auth")

	// Register the route for user registration
	api.POST("/register", userHandler.CreateUser)
}

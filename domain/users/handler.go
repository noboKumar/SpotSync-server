package users

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/noboKumar/SpotSync-server/domain/users/dto"
	"github.com/noboKumar/SpotSync-server/utils"
)

// handler is responsible for handling HTTP requests related to user operations.
type handler struct {
	service *service
}

// NewHandler creates a new instance of the user handler with the provided service.
func NewHandler(service *service) *handler {
	return &handler{service: service}
}

// CreateUser creates a new user. It validates the request payload, calls the service to create the user, and returns the appropriate HTTP response.
func (h *handler) CreateUser(c *echo.Context) error {
	var req dto.CreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusConflict, utils.Error{
			Code:    http.StatusConflict,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data:    response,
	})
}

func (h *handler) LoginUser(c *echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.LoginUser(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.Error{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "Login successful",
		Data:    response,
	})
}

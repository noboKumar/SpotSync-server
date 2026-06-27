package parkingzones

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/noboKumar/SpotSync-server/domain/parking_zones/dto"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateParkingZone(c *echo.Context) error {
	var req dto.CreateParkingZoneRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]any{
			"success": false,
			"message": "Invalid request payload",
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Validation failed",
			"details": err.Error(),
		})
	}

	response, err := h.service.CreateParkingZone(req)
	if err != nil {
		return c.JSON(500, map[string]any{
			"success": false,
			"message": "Internal server error",
			"details": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "Parking zone created successfully",
		Data:    response,
	})
}

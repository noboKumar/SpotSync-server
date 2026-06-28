package reservations

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/noboKumar/SpotSync-server/domain/reservations/dto"
	"github.com/noboKumar/SpotSync-server/utils"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service: service}
}

func (h *handler) ReserveParkingZone(c *echo.Context) error {
	var req dto.ReserveParkingZoneReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadGateway, utils.Error{
			Code:    http.StatusBadGateway,
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

	userID := c.Get("userID").(uint)
	req.UserID = userID

	res, err := h.service.ReserveParkingZone(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to reserve parking zone",
			Details: err.Error(),
		})
	}

	return c.JSON(201, dto.SuccessResponse{
		Success: true,
		Message: "Parking zone reserved successfully",
		Data:    res,
	})
}

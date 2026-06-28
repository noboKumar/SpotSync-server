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

	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(401, map[string]any{
			"success": false,
			"message": "Unauthorized",
		})
	}

	req.UserID = userID

	res, err := h.service.ReserveParkingZone(req, userID)
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

func (h *handler) GetMyReservation(c *echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"success": false,
			"message": "Unauthorized",
		})
	}

	reservations, err := h.service.GetMyReservation(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve reservations",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "Reservations retrieved successfully",
		Data:    reservations,
	})
}

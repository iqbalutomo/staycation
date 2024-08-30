package handler

import (
	"net/http"
	model "staycation/internal/models"
	service "staycation/internal/services"
	"staycation/pkg/utils"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type InvoiceHandler struct {
	Service service.InvoiceService
}

func NewInvoiceHandler(service service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{Service: service}
}

// ShowAccount godoc
// @Summary      Booking Room
// @Description  Book room from hotel
// @Tags         Booking
// @Accept       json
// @Produce      json
// @Param       booking body model.Booking true "Booking details"
// @Success      200  {object}  model.BookingSuccessResponse "Success"
// @Failure      400  {object}  utils.APIError
// @Failure      404  {object}  utils.APIError
// @Failure      500  {object}  utils.APIError
// @Security BearerAuth
// @Router       /bookings [post]
func (h *InvoiceHandler) BookRoom(c echo.Context) error {
	userClaims := c.Get("user").(jwt.MapClaims)
	userID := uint(userClaims["user_id"].(float64))
	email := userClaims["email"].(string)

	booking := new(model.Booking)
	if err := c.Bind(booking); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.BookingBadRequestErr, "invalid body request"))
	}

	if err := c.Validate(booking); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelValidationErr, err))
	}

	respData, err := h.Service.BookRoom(userID, email, booking)
	if err != nil {
		if err.Error() == "room_not_found" {
			return utils.HandleError(c, utils.NewNotFoundError(utils.BookingNotFound, "room not found"))
		}
		if err.Error() == "already-booked" {
			return utils.HandleError(c, utils.NewBadRequestError(utils.BookingBadRequestErr, "room is already booked for the requested dates"))
		}
		if err.Error() == "validate-time" {
			return utils.HandleError(c, utils.NewBadRequestError(utils.BookingBadRequestErr, "check-in must be at 14:00 and check-out must be at 12:00"))
		}
		if err.Error() == "booking-duration" {
			return utils.HandleError(c, utils.NewBadRequestError(utils.BookingBadRequestErr, "booking duration must be at least 1 full day"))
		}
		if err.Error() == "balance_not_found" {
			return utils.HandleError(c, utils.NewNotFoundError(utils.BookingNotFound, "balance not found"))
		}
		if err.Error() == "insufficient_balance" {
			return utils.HandleError(c, utils.NewBadRequestError(utils.BookingBadRequestErr, "insufficient balance"))
		}
		return utils.HandleError(c, utils.NewInternalError(utils.BookingInternalErr, err.Error()))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   respData,
	})
}

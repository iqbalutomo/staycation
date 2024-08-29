package webhook

import (
	"net/http"
	repository "staycation/internal/repositories"
	model "staycation/pkg/third_parties/xendit/models"
	"staycation/pkg/utils"

	"github.com/labstack/echo/v4"
)

type XenditWebhookHandler struct {
	InvoiceRepo repository.InvoiceRepository
	BalanceRepo repository.BalanceRepository
	HotelRepo   repository.HotelRepository
}

func NewXenditWebhookHandler(InvoiceRepo repository.InvoiceRepository, BalanceRepo repository.BalanceRepository, HotelRepo repository.HotelRepository) *XenditWebhookHandler {
	return &XenditWebhookHandler{InvoiceRepo, BalanceRepo, HotelRepo}
}

func (h *XenditWebhookHandler) InvoiceWebhook(c echo.Context) error {
	payload := new(model.XenditInvoiceWebhookPayload)
	if err := c.Bind(payload); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.InvoiceBadRequestErr, "invalid body request"))
	}

	// with id from invoice response (xendit)
	invoice, err := h.InvoiceRepo.FindInvoiceByID(payload.ID)
	if err != nil {
		return utils.HandleError(c, utils.NewInternalError(utils.InvoiceInternalErr, err.Error()))
	}

	if err := h.InvoiceRepo.UpdateInvoiceStatus(invoice.BookingID, "PAID"); err != nil {
		return utils.HandleError(c, utils.NewInternalError(utils.InvoiceInternalErr, err.Error()))
	}

	booking, err := h.InvoiceRepo.FindBookingByID(invoice.BookingID)
	if err != nil || booking == nil {
		return utils.HandleError(c, utils.NewNotFoundError(utils.BookingNotFound, "booking not found"))
	}

	balance, err := h.BalanceRepo.FindByUserID(booking.UserID)
	if err != nil || balance == nil {
		return utils.HandleError(c, utils.NewNotFoundError(utils.BalanceNotFound, "balance not found"))
	}

	if balance.Balance < booking.TotalPrice {
		return utils.HandleError(c, utils.NewBadRequestError(utils.BalanceInvalidReqErr, "insufficient balance"))
	}

	balance.Balance -= booking.TotalPrice
	if err := h.BalanceRepo.Update(balance); err != nil {
		return utils.HandleError(c, utils.NewInternalError(utils.BalanceInternalErr, err.Error()))
	}

	hotel, err := h.HotelRepo.FindByID(booking.HotelID)
	if err != nil || hotel == nil {
		return utils.HandleError(c, utils.NewNotFoundError(utils.HotelNotFound, "hotel not found"))
	}

	hotelBalance, err := h.BalanceRepo.FindByUserID(hotel.OwnerID)
	if err != nil || hotelBalance == nil {
		return utils.HandleError(c, utils.NewNotFoundError(utils.HotelNotFound, "hotel owner balance not found"))
	}

	hotelBalance.Balance += booking.TotalPrice
	if err := h.BalanceRepo.Update(hotelBalance); err != nil {
		return utils.HandleError(c, utils.NewInternalError(utils.BalanceInternalErr, err.Error()))
	}

	return c.String(http.StatusOK, "thanks mate to the paid.... [holy insane🥷🏼]")
}

package route

import (
	handler "staycation/internal/handlers"
	repository "staycation/internal/repositories"
	service "staycation/internal/services"
	"staycation/pkg/middlewares"
	"staycation/pkg/third_parties/xendit/webhook"

	"github.com/labstack/echo/v4"
)

func InvoiceRouter(e *echo.Echo) {
	roomRepo := repository.NewRoomRepository()
	roomTypeRepo := repository.NewRoomTypeRepository()
	balanceRepo := repository.NewBalanceRepository()
	hotelRepo := repository.NewHotelRepository()
	invoiceRepo := repository.NewInvoiceRepository()
	invoiceService := service.NewInvoiceService(invoiceRepo, roomRepo, roomTypeRepo, balanceRepo)
	invoiceHandler := handler.NewInvoiceHandler(invoiceService)

	xenditHandler := webhook.NewXenditWebhookHandler(invoiceRepo, balanceRepo, hotelRepo)

	e.POST("/bookings", invoiceHandler.BookRoom, middlewares.ProtectedRoute)
	e.POST("/invoice_webhook", xenditHandler.InvoiceWebhook)
}

package route

import (
	handler "staycation/internal/handlers"
	repository "staycation/internal/repositories"
	service "staycation/internal/services"
	"staycation/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

func InvoiceRouter(e *echo.Echo) {
	roomRepo := repository.NewRoomRepository()
	roomTypeRepo := repository.NewRoomTypeRepository()
	balanceRepo := repository.NewBalanceRepository()
	invoiceRepo := repository.NewInvoiceRepository()
	invoiceService := service.NewInvoiceService(invoiceRepo, roomRepo, roomTypeRepo, balanceRepo)
	invoiceHandler := handler.NewInvoiceHandler(invoiceService)

	e.POST("/bookings", invoiceHandler.BookRoom, middlewares.ProtectedRoute)
}

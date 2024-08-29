package route

import (
	handler "staycation/internal/handlers"
	repository "staycation/internal/repositories"
	service "staycation/internal/services"
	"staycation/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

func HotelRouter(e *echo.Echo) {
	hotelRepo := repository.NewHotelRepository()
	hotelService := service.NewHotelService(hotelRepo)
	hotelHandler := handler.NewHotelHandler(hotelService)

	e.POST("/hotel", hotelHandler.PostHotel, middlewares.ProtectedRoute)
}

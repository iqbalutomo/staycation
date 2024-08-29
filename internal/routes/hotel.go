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

	e.POST("/hotels", hotelHandler.PostHotel, middlewares.ProtectedRoute)
	e.PUT("/hotels/:id", hotelHandler.PutHotel, middlewares.ProtectedRoute)

	e.POST("/hotels/:hotel-id/roomtypes", hotelHandler.PostRoomType, middlewares.ProtectedRoute)
	e.POST("/hotels/:roomtype-id/room", hotelHandler.PostRoom, middlewares.ProtectedRoute)

	e.GET("/hotels", hotelHandler.GetHotels)
}

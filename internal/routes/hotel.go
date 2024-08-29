package route

import (
	handler "staycation/internal/handlers"
	repository "staycation/internal/repositories"
	service "staycation/internal/services"
	"staycation/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

func HotelRouter(e *echo.Echo) {
	// HOTEL (CRUD)
	hotelRepo := repository.NewHotelRepository()
	hotelService := service.NewHotelService(hotelRepo)
	hotelHandler := handler.NewHotelHandler(hotelService)

	e.POST("/hotels", hotelHandler.PostHotel, middlewares.ProtectedRoute)
	e.GET("/hotels", hotelHandler.GetHotels)
	e.PUT("/hotels/:id", hotelHandler.PutHotel, middlewares.ProtectedRoute)
	e.DELETE("/hotels/:id", hotelHandler.DeleteHotel, middlewares.ProtectedRoute)

	// ROOM TYPE
	roomTypeRepo := repository.NewRoomTypeRepository()
	roomTypeService := service.NewRoomTypeService(roomTypeRepo, hotelRepo)
	roomTypeHandler := handler.NewRoomTypeHandler(roomTypeService)

	e.POST("/hotels/:hotel-id/roomtypes", roomTypeHandler.PostRoomType, middlewares.ProtectedRoute)

	// ROOM
	roomRepo := repository.NewRoomRepository()
	roomService := service.NewRoomService(roomRepo, roomTypeRepo, hotelRepo)
	roomHandler := handler.NewRoomHandler(roomService)

	e.POST("/hotels/:roomtype-id/room", roomHandler.PostRoom, middlewares.ProtectedRoute)

}

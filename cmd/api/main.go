package main

import (
	"net/http"
	"staycation/config"
	repository "staycation/internal/repositories"
	route "staycation/internal/routes"
	service "staycation/internal/services"
	"staycation/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.InitConfig()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v := validator.New()
	e.Validator = &utils.CustomValidator{Validator: v}

	roomRepo := repository.NewRoomRepository()
	roomTypeRepo := repository.NewRoomTypeRepository()
	hotelRepo := repository.NewHotelRepository()
	roomService := service.NewRoomService(roomRepo, roomTypeRepo, hotelRepo)

	cronJobService := utils.NewCronJobService(roomService)
	cronJobService.Start()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Welcome staycation API bruh..",
		})
	})

	route.MainRouter(e)

	e.Logger.Fatal(e.Start(":" + config.Port))
}

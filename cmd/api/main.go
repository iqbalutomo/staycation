package main

import (
	"net/http"
	"staycation/config"
	repository "staycation/internal/repositories"
	route "staycation/internal/routes"
	service "staycation/internal/services"
	"staycation/pkg/utils"

	_ "staycation/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title           Staycation API
// @version         0.0.1
// @description     This is API for booking hotel with Xendit payment gateway
// @termsOfService  https://github.com/iqbalutomo/staycation

// @contact.name   API Support
// @contact.email  muhlisiqbalutomo@gmail.com

// @license.name  MIT License
// @license.url   https://github.com/iqbalutomo/staycation/blob/master/LICENSE

// @host      fc76-103-18-34-211.ngrok-free.app

// @securityDefinitions.apiKey  BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  Github
// @externalDocs.url          https://github.com/iqbalutomo
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
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	route.MainRouter(e)

	e.Logger.Fatal(e.Start(":" + config.Port))
}

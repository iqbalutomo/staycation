package main

import (
	"net/http"
	"staycation/config"
	route "staycation/internal/routes"
	"staycation/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v := validator.New()
	e.Validator = &utils.CustomValidator{Validator: v}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Welcome staycation API bruh..",
		})
	})

	route.MainRouter(e)

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}

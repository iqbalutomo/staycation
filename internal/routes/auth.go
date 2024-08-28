package route

import (
	handler "staycation/internal/handlers"
	repository "staycation/internal/repositories"
	service "staycation/internal/services"

	"github.com/labstack/echo/v4"
)

func AuthRouter(e *echo.Echo) {
	userRepo := repository.NewUserRepository()
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	e.POST("/users/register", authHandler.Register)
	e.POST("/users/login", authHandler.Login)
}

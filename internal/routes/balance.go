package route

import (
	handler "staycation/internal/handlers"
	repository "staycation/internal/repositories"
	service "staycation/internal/services"
	"staycation/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

func BalanceRouter(e *echo.Echo) {
	balanceRepo := repository.NewBalanceRepository()
	balanceService := service.NewBalanceService(balanceRepo)
	balanceHandler := handler.NewBalanceHandler(balanceService)

	e.POST("/users/deposit", balanceHandler.TopUp, middlewares.ProtectedRoute)
}

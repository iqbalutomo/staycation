package route

import "github.com/labstack/echo/v4"

func MainRouter(e *echo.Echo) {
	AuthRouter(e)
	BalanceRouter(e)
	HotelRouter(e)
}

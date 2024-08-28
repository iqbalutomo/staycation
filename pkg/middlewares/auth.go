package middlewares

import (
	"os"
	"staycation/pkg/utils"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// auth middleware
func ProtectedRoute(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return utils.HandleError(c, utils.NewUnauthorizedError(utils.Unauthorized, "unauthorized"))
		}

		// trim if any prefix like Bearer
		tokenString := strings.TrimPrefix(authHeader, "Bearer")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil || !token.Valid {
			return utils.HandleError(c, utils.NewUnauthorizedError(utils.InvalidCredential, "invalid credentials"))
		}

		c.Set("user", token.Claims.(jwt.MapClaims))
		return next(c)
	}
}

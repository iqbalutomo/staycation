package utils

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// check err is validator invalid
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid validation input")
		}

		errMap := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			switch err.Tag() {
			case "required":
				errMap[field] = "this field is required."
			case "email":
				errMap[field] = "invalid email format."
			case "oneof":
				errMap[field] = "invalid value provided."
			default:
				errMap[field] = "validation failed."
			}
		}

		return echo.NewHTTPError(http.StatusBadRequest, errMap)
	}

	return nil
}

func ValidatePhoneFormat(phone string) error {
	phoneRegex := `^(\+62|62|0)[1-9][0-9]{6,9}$`
	if !regexp.MustCompile(phoneRegex).MatchString(phone) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid phone number format.")
	}

	return nil
}

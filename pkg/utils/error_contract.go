package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	// REGISTER
	RegisterValidationErr = "REGISTER_VALIDATION_ERROR"
	RegisterEmailExist    = "REGISTER_EMAIL_EXIST"
	RegisteerPhoneExist   = "REGISTER_PHONE_EXIST"
	RegisterInternalErr   = "REGISTER_INTERNAL_ERROR"

	// LOGIN
	LoginValidationErr    = "LOGIN_VALIDATION_ERROR"
	LoginEmailPassInvalid = "LOGIN_INVALID_EMAIL_PASSWORD"
	LoginInternalErr      = "LOGIN_INTERNAL_ERROR"

	// AUTH
	Unauthorized      = "AUNAUTHORIZED"
	InvalidCredential = "INVALID_CREDENTIALS"

	// BALANCE
	BalanceInvalidReqErr = "DEPOSIT_BAD_REQUEST"
	BalanceInternalErr   = "DEPOSIT_INTERNAL_ERROR"
	BalanceMinTopUp      = "MINIMAL_TOPUP"
	BalanceMaxTopUp      = "MAXIMAL_TOPUP"
)

type APIError struct {
	Status  int         `json:"-"`
	Code    string      `json:"error_code"`
	Message interface{} `json:"message"`
	Detail  string      `json:"detail,omitempty"`
}

func NewNotFoundError(code, message string) *APIError {
	return &APIError{
		Status:  http.StatusNotFound,
		Code:    code,
		Message: message,
		Detail:  "Resource not found",
	}
}

func NewBadRequestError(code string, message interface{}) *APIError {
	return &APIError{
		Status:  http.StatusBadRequest,
		Code:    code,
		Message: message,
		Detail:  "Invalid request data",
	}
}

func NewInternalError(code, message string) *APIError {
	return &APIError{
		Status:  http.StatusInternalServerError,
		Code:    code,
		Message: message,
		Detail:  "Internal server error",
	}
}

func NewUnauthorizedError(code, message string) *APIError {
	return &APIError{
		Status:  http.StatusUnauthorized,
		Code:    code,
		Message: message,
		Detail:  "Unauthorized access",
	}
}

func HandleError(c echo.Context, err *APIError) error {
	return c.JSON(err.Status, err)
}

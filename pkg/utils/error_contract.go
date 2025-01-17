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
	BalanceNotFound      = "DEPOSIT_NOT_FOUND"
	BalanceInternalErr   = "DEPOSIT_INTERNAL_ERROR"
	BalanceMinTopUp      = "MINIMAL_TOPUP"
	BalanceMaxTopUp      = "MAXIMAL_TOPUP"

	// HOTEL
	HotelBadRequestErr = "HOTEL_BAD_REQUEST"
	HotelNotFound      = "HOTEL_NOT_FOUND"
	HotelValidationErr = "HOTEL_VALIDATION_ERROR"
	HotelInternalErr   = "HOTEL_INTERNAL_ERROR"
	HotelEmailExist    = "HOTEL_EMAIL_EXIST"
	HotelPhoneExist    = "HOTEL_PHONE_EXIST"

	// ROOM TYPE
	RoomTypeBadRequestErr = "ROOM_TYPE_BAD_REQUEST"
	RoomTypeNotFound      = "ROOM_TYPE_NOT_FOUND"
	RoomTypeValidationErr = "ROOM_TYPE_VALIDATION_ERROR"
	RoomTypeInternalErr   = "ROOM_TYPE_INTERNAL_ERROR"

	// ROOM
	RoomBadRequestErr = "ROOM_BAD_REQUEST"
	RoomNotFound      = "ROOM_NOT_FOUND"
	RoomValidationErr = "ROOM_VALIDATION_ERROR"
	RoomInternalErr   = "ROOM_INTERNAL_ERROR"

	// BOOKING
	BookingBadRequestErr = "BOOKING_BAD_REQUEST"
	BookingNotFound      = "BOOKING_NOT_FOUND"
	BookingValidationErr = "BOOKING_VALIDATION_ERROR"
	BookingInternalErr   = "BOOKING_INTERNAL_ERROR"

	// INVOICE
	InvoiceBadRequestErr = "INVOICE_BAD_REQUEST"
	InvoiceNotFound      = "INVOICE_NOT_FOUND"
	InvoiceValidationErr = "INVOICE_VALIDATION_ERROR"
	InvoiceInternalErr   = "INVOICE_INTERNAL_ERROR"

	// PAYMENT
	PaymentInternalErr = "PAYMENT_INTERNAL_ERROR"
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

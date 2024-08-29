package handler

import (
	"fmt"
	"net/http"
	model "staycation/internal/models"
	service "staycation/internal/services"
	"staycation/pkg/utils"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type hotelHandler struct {
	service service.HotelService
}

func NewHotelHandler(service service.HotelService) *hotelHandler {
	return &hotelHandler{service}
}

func (h *hotelHandler) PostHotel(c echo.Context) error {
	userClaims := c.Get("user").(jwt.MapClaims)
	userID := userClaims["user_id"].(float64)
	role := userClaims["role"].(string)

	if role != "hotel_owner" {
		return utils.HandleError(c, utils.NewUnauthorizedError(utils.InvalidCredential, "only user has owner hotel role"))
	}

	hotel := new(model.Hotel)
	if err := c.Bind(hotel); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "invalid body request"))
	}

	hotel.OwnerID = uint(userID)

	if err := c.Validate(hotel); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelValidationErr, err))
	}

	if err := utils.ValidatePhoneFormat(hotel.Phone); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelValidationErr, err))
	}

	createdHotel, err := h.service.CreateHotel(hotel)
	if err != nil {
		if err.Error() == "email_exist" {
			return utils.HandleError(c, utils.NewBadRequestError(utils.HotelEmailExist, "email already axist"))
		}
		if err.Error() == "phone_exist" {
			return utils.HandleError(c, utils.NewBadRequestError(utils.HotelPhoneExist, "phone number already exist"))
		}
		return utils.HandleError(c, utils.NewInternalError(utils.HotelInternalErr, err.Error()))
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status": "success",
		"data":   createdHotel,
	})
}

func (h *hotelHandler) GetHotels(c echo.Context) error {
	limitParam := c.QueryParam("limit")
	offsetParam := c.QueryParam("offset")
	limit := 10
	offset := 0

	if limitParam != "" {
		parsedLimit, err := strconv.Atoi(limitParam)
		if err != nil {
			return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "invalid limit query param"))
		}
		limit = parsedLimit
	}

	if offsetParam != "" {
		parsedOffset, err := strconv.Atoi(offsetParam)
		if err != nil {
			return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "invalid offset query param"))
		}
		offset = parsedOffset
	}

	respData, err := h.service.GetHotels(limit, offset)
	if err != nil {
		return utils.HandleError(c, utils.NewInternalError(utils.HotelInternalErr, err.Error()))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   respData,
	})
}

func (h *hotelHandler) PutHotel(c echo.Context) error {
	userClaims := c.Get("user").(jwt.MapClaims)
	userID := userClaims["user_id"].(float64)
	role := userClaims["role"].(string)

	if role != "hotel_owner" {
		return utils.HandleError(c, utils.NewUnauthorizedError(utils.InvalidCredential, "only user has owner hotel role"))
	}

	hotelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "invalid hotel id param"))
	}

	hotel, err := h.service.FindHotelByID(uint(hotelID))
	if err != nil {
		if err.Error() == "hotel_not_found" {
			return utils.HandleError(c, utils.NewNotFoundError(utils.HotelNotFound, "hotel not found"))
		}
		return utils.HandleError(c, utils.NewInternalError(utils.HotelInternalErr, err.Error()))
	}

	if err := c.Bind(hotel); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "invalid body request"))
	}

	if err := c.Validate(hotel); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelValidationErr, err))
	}

	if err := utils.ValidatePhoneFormat(hotel.Phone); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelValidationErr, err))
	}

	updatedHotel, err := h.service.UpdateHotel(userID, hotel)
	if err != nil {
		if err.Error() == "invalid_credentials" {
			return utils.HandleError(c, utils.NewUnauthorizedError(utils.Unauthorized, "invalid credentials"))
		}
		if err.Error() == "hotel_not_found" {
			return utils.HandleError(c, utils.NewNotFoundError(utils.HotelNotFound, "hotel not found"))
		}
		return utils.HandleError(c, utils.NewInternalError(utils.HotelInternalErr, err.Error()))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   updatedHotel,
	})

}

func (h *hotelHandler) DeleteHotel(c echo.Context) error {
	userClaims := c.Get("user").(jwt.MapClaims)
	userID := userClaims["user_id"].(float64)

	hotelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "invalid hotel id param"))
	}

	hotel, err := h.service.FindHotelByID(uint(hotelID))
	if err != nil {
		return utils.HandleError(c, utils.NewInternalError(utils.HotelInternalErr, err.Error()))
	} else if hotel == nil {
		return utils.HandleError(c, utils.NewNotFoundError(utils.HotelNotFound, "hotel not found"))
	}

	if hotel.OwnerID != uint(userID) {
		return utils.HandleError(c, utils.NewUnauthorizedError(utils.Unauthorized, "invalid credentials"))
	}

	if err := h.service.DeleteHotel(hotelID); err != nil {
		return utils.HandleError(c, utils.NewInternalError(utils.HotelInternalErr, err.Error()))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "success",
		"message": fmt.Sprintf("hotel with id %d has been deleted temporary", hotel.ID),
	})
}

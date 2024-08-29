package handler

import (
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

	hotel := new(model.Hotel)
	if err := c.Bind(hotel); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "invalid request"))
	}

	hotel.OwnerID = uint(userID)

	if err := c.Validate(hotel); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelValidationErr, err))
	}

	if role != "hotel_owner" {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "only user has owner hotel role"))
	}

	createdHotel, err := h.service.NewHotel(hotel)
	if err != nil {
		if err.Error() == "email_exist" {
			return utils.HandleError(c, utils.NewBadRequestError(utils.HotelEmailExist, "email already axist."))
		}
		if err.Error() == "phone_exist" {
			return utils.HandleError(c, utils.NewBadRequestError(utils.HotelPhoneExist, "phone number already exist."))
		}
		return utils.HandleError(c, utils.NewInternalError(utils.HotelInternalErr, err.Error()))
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status": "success",
		"data":   createdHotel,
	})
}

func (h *hotelHandler) PostRoomType(c echo.Context) error {
	userClaims := c.Get("user").(jwt.MapClaims)
	userID := userClaims["user_id"].(float64)
	role := userClaims["role"].(string)

	reqBody := new(model.RoomTypeRequest)

	hotelID, err := strconv.Atoi(c.Param("hotel-id"))
	if err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "invalid request"))
	}

	if err := c.Bind(reqBody); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "invalid request"))
	}

	reqBody.RoomType.HotelID = uint(hotelID)

	if err := c.Validate(reqBody); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelValidationErr, err))
	}

	if role != "hotel_owner" {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "only user has owner hotel role"))
	}

	createdRoomType, err := h.service.NewRoomType(userID, &reqBody.RoomType, &reqBody.RoomBedType, &reqBody.RoomFacilities)
	if err != nil {
		return utils.HandleError(c, utils.NewInternalError(utils.HotelInternalErr, err.Error()))
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status": "success",
		"data":   createdRoomType,
	})
}

func (h *hotelHandler) PostRoom(c echo.Context) error {
	userClaims := c.Get("user").(jwt.MapClaims)
	userID := userClaims["user_id"].(float64)
	role := userClaims["role"].(string)

	reqBody := new(model.Room)

	roomTypeID, err := strconv.Atoi(c.Param("roomtype-id"))
	if err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "invalid request"))
	}

	if err := c.Bind(reqBody); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "invalid request"))
	}

	reqBody.RoomTypeID = uint(roomTypeID)

	if err := c.Validate(reqBody); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelValidationErr, err))
	}

	if role != "hotel_owner" {
		return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "only user has owner hotel role"))
	}

	createdRoom, err := h.service.NewRoom(userID, reqBody)
	if err != nil {
		return utils.HandleError(c, utils.NewInternalError(utils.HotelInternalErr, err.Error()))
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status": "success",
		"data":   createdRoom,
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
			return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "invalid request"))
		}
		limit = parsedLimit
	}

	if offsetParam != "" {
		parsedOffset, err := strconv.Atoi(offsetParam)
		if err != nil {
			return utils.HandleError(c, utils.NewBadRequestError(utils.HotelBadRequestErr, "invalid request"))
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

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

type roomTypeHandler struct {
	service service.RoomTypeService
}

func NewRoomTypeHandler(service service.RoomTypeService) *roomTypeHandler {
	return &roomTypeHandler{service}
}

func (h *roomTypeHandler) PostRoomType(c echo.Context) error {
	userClaims := c.Get("user").(jwt.MapClaims)
	userID := userClaims["user_id"].(float64)
	role := userClaims["role"].(string)

	if role != "hotel_owner" {
		return utils.HandleError(c, utils.NewUnauthorizedError(utils.InvalidCredential, "only user has owner hotel role"))
	}

	hotelID, err := strconv.Atoi(c.Param("hotel-id"))
	if err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.RoomTypeBadRequestErr, "invalid hotel id param"))
	}

	reqBody := new(model.RoomTypeRequest)
	if err := c.Bind(reqBody); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.RoomTypeBadRequestErr, "invalid body request"))
	}

	reqBody.RoomType.HotelID = uint(hotelID)

	if err := c.Validate(reqBody); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.RoomTypeValidationErr, err))
	}

	createdRoomType, err := h.service.CreateRoomType(userID, &reqBody.RoomType, &reqBody.RoomBedType, &reqBody.RoomFacilities)
	if err != nil {
		if err.Error() == "hotel_not_found" {
			return utils.HandleError(c, utils.NewNotFoundError(utils.RoomTypeNotFound, "hotel not found"))
		}
		if err.Error() == "invalid_credentials" {
			return utils.HandleError(c, utils.NewUnauthorizedError(utils.Unauthorized, "invalid credentials"))
		}
		return utils.HandleError(c, utils.NewInternalError(utils.HotelInternalErr, err.Error()))
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status": "success",
		"data":   createdRoomType,
	})
}

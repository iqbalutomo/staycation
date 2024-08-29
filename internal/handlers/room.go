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

type roomHandler struct {
	service service.RoomService
}

func NewRoomHandler(service service.RoomService) *roomHandler {
	return &roomHandler{service}
}

func (h *roomHandler) PostRoom(c echo.Context) error {
	userClaims := c.Get("user").(jwt.MapClaims)
	userID := userClaims["user_id"].(float64)
	role := userClaims["role"].(string)

	if role != "hotel_owner" {
		return utils.HandleError(c, utils.NewUnauthorizedError(utils.InvalidCredential, "only user has owner hotel role"))
	}

	roomTypeID, err := strconv.Atoi(c.Param("roomtype-id"))
	if err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.RoomBadRequestErr, "invalid roomtype id param"))
	}

	reqBody := new(model.Room)
	if err := c.Bind(reqBody); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.RoomBadRequestErr, "invalid body request"))
	}

	reqBody.RoomTypeID = uint(roomTypeID)

	if err := c.Validate(reqBody); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.RoomValidationErr, err))
	}

	createdRoom, err := h.service.CreateRoom(userID, reqBody)
	if err != nil {
		if err.Error() == "roomtype_not_found" {
			return utils.HandleError(c, utils.NewNotFoundError(utils.RoomNotFound, "room type not found"))
		}
		if err.Error() == "invalid_credentials" {
			return utils.HandleError(c, utils.NewUnauthorizedError(utils.InvalidCredential, "invalid credentials"))
		}
		return utils.HandleError(c, utils.NewInternalError(utils.RoomInternalErr, err.Error()))
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status": "success",
		"data":   createdRoom,
	})
}

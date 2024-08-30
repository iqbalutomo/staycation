package handler

import (
	"net/http"
	model "staycation/internal/models"
	service "staycation/internal/services"

	email_mailtrap "staycation/pkg/third_parties/mailtrap/email"
	"staycation/pkg/utils"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *authHandler {
	return &authHandler{service: service}
}

func (h *authHandler) Register(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	phone := c.FormValue("phone")
	role := c.FormValue("role")

	formData := model.User{
		Name:     name,
		Email:    email,
		Password: password,
		Phone:    phone,
		Role:     model.UserRoleEnum(role),
	}

	if err := c.Validate(&formData); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.RegisterValidationErr, err))
	}

	// validating phone number format
	if err := utils.ValidatePhoneFormat(formData.Phone); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.RegisterValidationErr, err))
	}

	user, err := h.service.Register(formData)
	if err != nil {
		if err.Error() == "email_exist" {
			return utils.HandleError(c, utils.NewBadRequestError(utils.RegisterEmailExist, "email already axist."))
		}
		if err.Error() == "phone_exist" {
			return utils.HandleError(c, utils.NewBadRequestError(utils.RegisteerPhoneExist, "phone number already exist."))
		}
		return utils.HandleError(c, utils.NewInternalError(utils.RegisterInternalErr, err.Error()))
	}

	if err := email_mailtrap.SendEmailRegister(formData.Email, formData.Name); err != nil {
		return utils.HandleError(c, utils.NewInternalError(utils.RegisterInternalErr, err.Error()))
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status": "success",
		"data":   user,
	})
}

func (h *authHandler) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	formData := model.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	if err := c.Validate(&formData); err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.LoginValidationErr, err.Error()))
	}

	resp, err := h.service.Login(formData)
	if err != nil {
		if err.Error() == "invalid_emailorpassword" {
			return utils.HandleError(c, utils.NewBadRequestError(utils.LoginEmailPassInvalid, "invalid email or password"))
		}
		return utils.HandleError(c, utils.NewInternalError(utils.LoginInternalErr, err.Error()))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   resp,
	})
}

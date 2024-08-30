package handler

import (
	"net/http"
	service "staycation/internal/services"
	"staycation/pkg/utils"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type balanceHandler struct {
	service service.BalanceService
}

func NewBalanceHandler(service service.BalanceService) *balanceHandler {
	return &balanceHandler{service}
}

// ShowAccount godoc
// @Summary      Top Up Balance
// @Description  Deposit your balance for book room from hotel
// @Tags         User
// @Accept       json
// @Produce      json
// @Param       amount query float64 true "Deposit amount"
// @Success      200  {object}  map[string]interface{} "Success"
// @Failure      400  {object}  utils.APIError
// @Failure      500  {object}  utils.APIError
// @Security BearerAuth
// @Router       /users/deposit [post]
func (h *balanceHandler) TopUp(c echo.Context) error {
	userClaims := c.Get("user").(jwt.MapClaims)
	userID := userClaims["user_id"].(float64)

	amount, err := strconv.ParseFloat(c.QueryParam("amount"), 64)
	if err != nil {
		return utils.HandleError(c, utils.NewBadRequestError(utils.BalanceInvalidReqErr, "invalid amount"))
	}

	if err := h.service.Deposit(int(userID), amount); err != nil {
		if err.Error() == "min_topup" {
			return utils.HandleError(c, utils.NewBadRequestError(utils.BalanceMinTopUp, "minimal topup Rp. 100.000"))
		} else if err.Error() == "max_topup" {
			return utils.HandleError(c, utils.NewBadRequestError(utils.BalanceMaxTopUp, "maximal topup Rp. 10.000.000"))
		}
		return utils.HandleError(c, utils.NewInternalError(utils.BalanceInternalErr, err.Error()))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "success",
		"message": "deposit successfully!",
	})
}

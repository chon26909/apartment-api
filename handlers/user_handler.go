package handlers

import (
	"net/http"

	"github.com/chon26909/apartment-api/models"
	"github.com/chon26909/apartment-api/services"
	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	GetUserList(ctx echo.Context) error
	GetProfile(ctx echo.Context) error
}

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) IUserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUserList(ctx echo.Context) error {

	users := []models.User{}

	return ctx.JSONPretty(http.StatusOK, users, "")
}

func (h *UserHandler) GetProfile(c echo.Context) error {

	ctx := c.Request().Context()

	userId := c.Get("userId").(int)

	res, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSONPretty(http.StatusOK, res, "")
}

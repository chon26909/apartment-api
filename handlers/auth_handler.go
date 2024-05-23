package handlers

import (
	"net/http"

	"github.com/chon26909/apartment-api/models"
	"github.com/chon26909/apartment-api/services"
	"github.com/labstack/echo/v4"
)

type IAuthHandler interface {
	Login(ctx echo.Context) error
	// Register(ctx echo.Echo) error
}

type AuthHandler struct {
	authService services.IAuthService
}

func NewAuthHandler(authService services.IAuthService) IAuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c echo.Context) error {

	ctx := c.Request().Context()

	request := new(models.LoginRequest)
	if err := c.Bind(&request); err != nil {
		return nil
	}

	res, _ := h.authService.Login(ctx, request)

	return c.JSONPretty(http.StatusOK, res, "")
}

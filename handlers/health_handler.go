package handlers

import (
	"net/http"

	"github.com/chon26909/apartment-api/services"
	"github.com/labstack/echo/v4"
)

type IHealthHandler interface {
	HealthCheck(ctx echo.Context) error
}

type healthHandler struct {
	healthService services.IHealthService
}

func NewHealthHandler(healthService services.IHealthService) IHealthHandler {
	return &healthHandler{healthService: healthService}
}

func (h *healthHandler) HealthCheck(ctx echo.Context) error {
	response, err := h.healthService.HealthCheck()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

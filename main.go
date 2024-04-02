package main

import (
	"fmt"

	"github.com/chon26909/apartment-api/config"
	"github.com/chon26909/apartment-api/handlers"
	"github.com/chon26909/apartment-api/lib"
	"github.com/chon26909/apartment-api/services"
	"github.com/labstack/echo/v4"
)

func main() {

	config := config.InitConfig()

	fmt.Printf("config %v", config)

	db := lib.NewMySqlConnection(config)

	_ = db

	healthService := services.NewHealthService()
	healthHandler := handlers.NewHealthHandler(healthService)

	// create router
	e := echo.New()
	e.GET("/health", healthHandler.HealthCheck)

	e.Start(fmt.Sprintf(":%d", config.App.Port))
}

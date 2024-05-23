package main

import (
	"fmt"

	"github.com/chon26909/apartment-api/config"
	"github.com/chon26909/apartment-api/handlers"
	"github.com/chon26909/apartment-api/lib"
	"github.com/chon26909/apartment-api/middleware"
	"github.com/chon26909/apartment-api/models"
	"github.com/chon26909/apartment-api/repositories"
	"github.com/chon26909/apartment-api/services"
	"github.com/labstack/echo/v4"
)

func main() {

	config := config.InitConfig()

	fmt.Printf("config %v", config)

	db := lib.NewMySqlConnection(config)

	db.AutoMigrate(models.User{})

	// repository
	userRepository := repositories.NewUserRepository(db)

	// service
	healthService := services.NewHealthService()
	authService := services.NewAuthService(userRepository)
	userService := services.NewUserService(userRepository)

	// handler
	healthHandler := handlers.NewHealthHandler(healthService)
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	// create router
	e := echo.New()
	e.GET("health", healthHandler.HealthCheck)

	v1 := e.Group("v1")

	auth := v1.Group("/auth")
	auth.POST("/login", authHandler.Login)

	user := v1.Group("/user", middleware.VerifyHeader())
	user.GET("", userHandler.GetUserList)
	user.GET("/me", userHandler.GetProfile)

	e.Start(fmt.Sprintf(":%d", config.App.Port))
}

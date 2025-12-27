package main

import (
	"go-rest-api/internal/config"
	"go-rest-api/internal/database"
	"go-rest-api/internal/handler"
	"go-rest-api/internal/repository"
	"go-rest-api/internal/router"
	"go-rest-api/internal/service"
)

func main() {

	cfg := config.Load()

	db := database.InitDB(cfg.DBUrl)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := router.New(userHandler)
	r.Run("localhost:" + cfg.AppPort)
}

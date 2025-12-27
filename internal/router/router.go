package router

import (
	"go-rest-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func New(
	userHandler *handler.UserHandler,
)  *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	api := r.Group("/api")

	RegisterUserRoutes(api, userHandler)

	return r
}
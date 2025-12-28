package router

import (
	"go-rest-api/internal/handler"
	"go-rest-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func New(
	userHandler *handler.UserHandler,
) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Apply Global Middleware (CORS)
	r.Use(middleware.CORSMiddleware())

	api := r.Group("/api")

	// Public routes
	api.POST("/login", userHandler.Login)
	RegisterPublicUserRoutes(api, userHandler)

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.JWTAuth())
	{
		RegisterProtectedUserRoutes(protected, userHandler)
	}

	return r
}

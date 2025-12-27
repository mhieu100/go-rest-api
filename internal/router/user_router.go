package router

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/internal/handler"
)

func RegisterUserRoutes(r *gin.RouterGroup, h *handler.UserHandler) {
	users := r.Group("/users")
	{
		users.GET("/:id", h.Get)
		users.POST("", h.Create)
	}
}

package router

import (
	"go-rest-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, h *handler.UserHandler) {
	users := r.Group("/users")
	{
		users.GET("", h.GetAll)
		users.GET("/:id", h.Get)
		users.POST("", h.Create)
		users.PUT("/:id", h.Update)
		users.DELETE("/:id", h.Delete)
	}
}

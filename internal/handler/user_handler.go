package handler

import (
	"go-rest-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
    service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
    return &UserHandler{service: s}
}

func (h *UserHandler) Get(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    user, err := h.service.GetUser(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Create(c *gin.Context) {
    var req struct {
        Email string `json:"email"`
        Name  string `json:"name"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := h.service.CreateUser(req.Email, req.Name)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusCreated)
}

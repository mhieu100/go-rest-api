package handler

import (
	"go-rest-api/internal/dto"
	"go-rest-api/internal/service"
	"go-rest-api/pkg/response"
	pkgValidator "go-rest-api/pkg/validator"
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
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}
	response.Success(c, user)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	users, total, err := h.service.GetAllUsers(page, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessWithPagination(c, users, total, page, limit)
}

func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := pkgValidator.FormatErrors(err)
		if validationErrors != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Validation failed",
				"errors":  validationErrors,
			})
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.CreateUser(req.Email, req.Name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Created(c, nil)
}

func (h *UserHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req dto.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := pkgValidator.FormatErrors(err)
		if validationErrors != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Validation failed",
				"errors":  validationErrors,
			})
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.UpdateUser(uint(id), req.Email, req.Name)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteUser(uint(id))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := pkgValidator.FormatErrors(err)
		if validationErrors != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Validation failed",
				"errors":  validationErrors,
			})
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Login(req.Email)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	response.Success(c, dto.LoginResponse{Token: token})
}

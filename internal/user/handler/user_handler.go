package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romulosm/go-modular-backend-template/internal/user/domain"
	"github.com/romulosm/go-modular-backend-template/internal/user/service"
	"github.com/romulosm/go-modular-backend-template/pkg/logger"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Log.Errorf("Erro ao fazer bind do JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.CreateUser(&user)
	if err != nil {
		logger.Log.Errorf("Erro ao criar usuário: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Log.Infof("Usuário criado com sucesso: %s", user.ID)
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userService.GetUser(id)
	if err != nil {
		logger.Log.Errorf("Erro ao buscar usuário: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	logger.Log.Infof("Usuário encontrado: %s", user.ID)
	c.JSON(http.StatusOK, user)
}

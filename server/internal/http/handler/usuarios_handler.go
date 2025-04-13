package handler

import (
	"net/http"
	"server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(us service.UserService) *UserHandler {
	return &UserHandler{us}
}

func (h UserHandler) GetUsuarios(c *gin.Context) {
	res, err := h.service.GetUsuarios()

	if err != nil {
		c.JSON(404, err)
	}
	c.JSON(http.StatusOK, res)
}
func (h UserHandler) GetUsuario(c *gin.Context) {
	id := c.Param("id")

	convertId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"message": "Erro ao converter id"})
	}

	res, err := h.service.GetUsuario(convertId)

	if err != nil {
		c.JSON(404, err)
	}
	c.JSON(http.StatusOK, res)
}

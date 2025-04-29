package handler

import (
	"fmt"
	"net/http"
	"server/internal/models"
	"server/internal/service"
	"server/internal/utils"
	"strconv"
	"strings"

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
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, models.HttpResponse{Status: http.StatusNotFound, Message: "O banco ainda não possui usuários"})
		return
	}
	c.JSON(http.StatusOK, models.HttpResponse{Status: http.StatusOK, Message: "Todos os usuarios do banco", Data: res})
}

func (h UserHandler) GetUsuario(c *gin.Context) {
	id := c.Param("id")

	res, err := h.service.GetUsuario(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.HttpResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	if res == nil {
		c.JSON(http.StatusNotFound, models.HttpResponse{Status: http.StatusNotFound, Message: "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, models.HttpResponse{Status: http.StatusOK, Message: fmt.Sprintf("Usuário de id %s encontrado", id), Data: res})
}

func (h UserHandler) DeleteUsuario(c *gin.Context) {
	id := c.Param("id")

	convertId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, models.HttpResponse{Status: http.StatusBadGateway, Message: "Erro ao converter número de id para inteiro"})
	}

	err = h.service.DeleteUsuario(convertId)

	if err != nil {
		if strings.Contains(err.Error(), "usuário não encontrado") {
			c.JSON(http.StatusNotFound, models.HttpResponse{Status: http.StatusNotFound, Message: err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, models.HttpResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.HttpResponse{Status: http.StatusOK, Message: "Usuário deletado com sucesso"})
}

func (h UserHandler) CreateUsuario(c *gin.Context) {
	var usuario models.Usuario

	err := c.BindJSON(&usuario)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HttpResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	err = h.service.CreateUsuario(usuario)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HttpResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, models.HttpResponse{Status: http.StatusCreated, Message: "Usuário criado com sucesso"})
}
func (h UserHandler) Login(c *gin.Context) {
	var usuario models.LoginUsuario

	if err := c.BindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, models.HttpResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	u, err := h.service.Login(usuario)

	if err != nil {
		if strings.Contains(err.Error(), "usuário não encontrado") {
			c.JSON(http.StatusUnauthorized, models.HttpResponse{Status: http.StatusUnauthorized, Message: "Usuário não existe"})
			return
		}
		if strings.Contains(err.Error(), "senha incorreta") {
			c.JSON(http.StatusUnauthorized, models.HttpResponse{Status: http.StatusUnauthorized, Message: "Senha incorreta"})
			return
		}
		c.JSON(http.StatusUnauthorized, models.HttpResponse{Status: http.StatusUnauthorized, Message: "Usuário ou senha incorreta"})
		return
	}

	token, err := utils.GenerateJwtToken(u.ID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HttpResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, models.HttpResponse{Status: http.StatusOK, Message: "Usuário logado com sucesso"})
}

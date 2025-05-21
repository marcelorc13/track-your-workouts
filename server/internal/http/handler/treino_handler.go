package handler

import (
	"fmt"
	"net/http"
	"server/internal/models"
	"server/internal/service"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type TreinoHandler struct {
	service service.TreinoService
}

func NewTreinoHandler(ts service.TreinoService) *TreinoHandler {
	return &TreinoHandler{ts}
}
func (th TreinoHandler) CreateTreino(c *gin.Context) {
	var treino models.Treino

	err := c.BindJSON(&treino)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HttpResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.HttpResponse{Status: http.StatusUnauthorized, Message: "Não possui token de autorização"})
		return
	}

	claims, err := utils.GetTokenClaims(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.HttpResponse{Status: http.StatusUnauthorized, Message: "Token mal formado"})
		return
	}

	userId, ok := claims["id"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, models.HttpResponse{Status: http.StatusUnauthorized, Message: "Token mal formado"})
		return
	}

	treino.CriadoPor = userId

	err = th.service.CreateTreino(treino)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.HttpResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, models.HttpResponse{Status: http.StatusCreated, Message: "Treino criado com sucesso"})
}

func (th TreinoHandler) GetTreinosDoUsuario(c *gin.Context) {
	usuarioId := c.Param("id")

	res, err := th.service.GetTreinosDoUsuario(usuarioId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HttpResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.HttpResponse{Status: http.StatusOK, Message: fmt.Sprintf("Treinos do usuário %s", usuarioId), Data: res})
}

func (th TreinoHandler) CreateSecao(c *gin.Context) {
	var secao models.Secao

	err := c.BindJSON(&secao)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HttpResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.HttpResponse{Status: http.StatusUnauthorized, Message: "Não possui token de autorização"})
		return
	}

	claims, err := utils.GetTokenClaims(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.HttpResponse{Status: http.StatusUnauthorized, Message: "Token mal formado"})
		return
	}

	userId, ok := claims["id"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, models.HttpResponse{Status: http.StatusUnauthorized, Message: "Token mal formado"})
		return
	}

	secao.IDUsuario = userId

	err = th.service.CreateSecao(secao)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.HttpResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, models.HttpResponse{Status: http.StatusCreated, Message: "Seção de treino criada com sucesso"})
}

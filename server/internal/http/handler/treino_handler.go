package handler

import (
	"net/http"
	"server/internal/models"
	"server/internal/service"

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

	err = th.service.CreateTreino(treino)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.HttpResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, models.HttpResponse{Status: http.StatusCreated, Message: "Treino criado com sucesso"})
}

func (th TreinoHandler) GetTreinos(c *gin.Context) {
	res, err := th.service.GetTreinos()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HttpResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.HttpResponse{Status: http.StatusOK, Message: "Todos os treinos do banco", Data: res})
}

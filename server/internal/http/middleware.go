package httpServer

import (
	"net/http"
	"server/internal/models"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.HttpResponse{Status: http.StatusUnauthorized, Message: "Você não tem acesso"})
		return
	}
	err = utils.VerifyJwtToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.HttpResponse{Status: http.StatusUnauthorized, Message: "Token inválido"})
	}
	c.Next()
}

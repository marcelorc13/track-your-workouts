package httpServer

import (
	"net/http"
	"server/internal/models"

	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	_, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.HttpResponse{Status: http.StatusUnauthorized, Message: "Você não tem acesso"})
		return
	}
	c.Next()
}

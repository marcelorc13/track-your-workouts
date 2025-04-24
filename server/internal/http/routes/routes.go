package routes

import (
	"server/internal/http/handler"

	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	router *gin.Engine
}

func NewRouteHandler(r *gin.Engine) *RouteHandler {
	return &RouteHandler{r}
}

func (rh RouteHandler) UserRoutes(uh handler.UserHandler) {
	rh.router.GET("usuarios/", uh.GetUsuarios)
	rh.router.GET("usuarios/:id", uh.GetUsuario)
	rh.router.DELETE("usuarios/:id", uh.DeleteUsuario)
	rh.router.POST("usuarios/", uh.CreateUsuario)
	rh.router.POST("usuarios/login", uh.Login)
}

func (rh RouteHandler) TreinoRoutes(th handler.TreinoHandler) {
	rh.router.POST("treinos/", th.CreateTreino)
}

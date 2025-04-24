package routes

import (
	"server/internal/http/handler"

	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	handler handler.UserHandler
	router  *gin.Engine
}

func NewRouteHandler(uh handler.UserHandler, r *gin.Engine) *RouteHandler {
	return &RouteHandler{uh, r}
}

func (rh RouteHandler) UserRoutes() {
	rh.router.GET("usuarios/", rh.handler.GetUsuarios)
	rh.router.GET("usuarios/:id", rh.handler.GetUsuario)
	rh.router.DELETE("usuarios/:id", rh.handler.DeleteUsuario)
	rh.router.POST("usuarios/", rh.handler.CreateUsuario)
	rh.router.POST("usuarios/login", rh.handler.Login)
}

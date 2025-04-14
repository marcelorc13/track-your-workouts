package routes

import (
	"server/internal/http/handler"

	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	handler handler.UserHandler
}

func NewRouteHandler(uh handler.UserHandler) *RouteHandler {
	return &RouteHandler{uh}
}

func (rh RouteHandler) UserRoutes(r *gin.Engine) {
	r.GET("usuarios/", rh.handler.GetUsuarios)
	r.GET("usuarios/:id", rh.handler.GetUsuario)
	r.DELETE("usuarios/:id", rh.handler.DeleteUsuario)
	r.POST("usuarios/", rh.handler.CreateUsuario)
}

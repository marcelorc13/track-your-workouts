package routes

import (
	httpServer "server/internal/http"
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
	usuarios := rh.router.Group("usuarios/")

	usuarios.GET("/", uh.GetUsuarios)
	usuarios.GET("/:id", uh.GetUsuario)
	usuarios.DELETE("/:id", uh.DeleteUsuario)
	usuarios.POST("/", uh.CreateUsuario)
	usuarios.POST("/login", uh.Login)
}

func (rh RouteHandler) TreinoRoutes(th handler.TreinoHandler) {
	treinos := rh.router.Group("treinos/")

	treinos.Use(httpServer.Middleware)

	treinos.POST("/", th.CreateTreino)
	treinos.GET("/", th.GetTreinos)
}

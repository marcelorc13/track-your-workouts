package httpServer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(r *gin.Engine) *Router {
	return &Router{r}
}

func (r Router) Start() {
	http.ListenAndServe(":8080", r.router)
}

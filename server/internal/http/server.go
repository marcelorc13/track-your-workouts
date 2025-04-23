package httpServer

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(r *gin.Engine) *Router {
	return &Router{r}
}

func (r Router) Start() {
	godotenv.Load()
	port := os.Getenv("PORT")
	http.ListenAndServe(port, r.router)
}

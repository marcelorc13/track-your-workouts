package main

import (
	httpServer "server/internal/http"
	"server/internal/http/handler"
	"server/internal/http/routes"
	"server/internal/repository"
	"server/internal/service"
	"server/pkg/database"

	"github.com/gin-contrib/cors"
)

func main() {
	router := routes.GetRouter()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

	router.Use(cors.New(config))

	r := httpServer.NewRouter(router)

	mysqlDB := database.ConnectMySQL()
	mongoDB := database.ConnectMongo()

	defer mysqlDB.Close()

	routHand := routes.NewRouteHandler(router)

	userRepo := repository.NewUserRepository(mysqlDB)
	userServ := service.NewUserService(*userRepo)
	userHand := handler.NewUserHandler(*userServ)
	routHand.UserRoutes(*userHand)

	treinoRepo := repository.NewTreinoRepository(mongoDB)
	treinoServ := service.NewTreinoService(*treinoRepo)
	treinoHand := handler.NewTreinoHandler(*treinoServ)
	routHand.TreinoRoutes(*treinoHand)

	r.Start()
}

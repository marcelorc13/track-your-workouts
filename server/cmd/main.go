package main

import (
	httpServer "server/internal/http"
	"server/internal/http/handler"
	"server/internal/http/routes"
	"server/internal/repository"
	"server/internal/service"
	"server/pkg/database"
)

func main() {
	router := routes.GetRouter()

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

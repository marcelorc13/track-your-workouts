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
	// mongoDB := database.ConnectMongo()

	defer mysqlDB.Close()

	userRepo := repository.NewUserRepository(mysqlDB)
	userServ := service.NewUserService(*userRepo)
	userHand := handler.NewUserHandler(*userServ)
	routHand := routes.NewRouteHandler(*userHand, router)

	// treinoRepo := repository.NewTreinoRepository(mongoDB)

	routHand.UserRoutes()

	r.Start()
}

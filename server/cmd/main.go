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

	db := database.ConnectMySQL()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userServ := service.NewUserService(*userRepo)
	userHand := handler.NewUserHandler(*userServ)
	routHand := routes.NewRouteHandler(*userHand)

	routHand.UserRoutes(router)

	r.Start()
}

package main

import (
	"event-booking/config"
	"event-booking/database"
	"event-booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnvVariables()
	database.InitDatabase()
	server := gin.Default()

	routes.RegisterUserRoutes(server)
	routes.RegisterEventRoutes(server)
	routes.RegisterRegistrationRoutes(server)

	server.Run(":3000")
}

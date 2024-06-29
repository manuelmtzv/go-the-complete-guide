package main

import (
	"event-booking/database"
	"event-booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()
	server := gin.Default()

	routes.RegisterEventRoutes(server)
	routes.RegisterUserRoutes(server)

	server.Run(":3000")
}

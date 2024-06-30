package routes

import (
	"event-booking/controllers"
	"event-booking/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRegistrationRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events/:id/register", controllers.RegisterForEvent)
}

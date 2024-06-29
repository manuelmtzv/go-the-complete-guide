package routes

import (
	"event-booking/controllers"
	"event-booking/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEventById)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", controllers.CreateEvent)
	authenticated.PATCH("/events/:id", controllers.UpdateEvent)
	authenticated.DELETE("/events/:id", controllers.DeleteEvent)
}

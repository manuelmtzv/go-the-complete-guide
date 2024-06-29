package routes

import (
	"event-booking/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEventById)
	server.POST("/events", controllers.CreateEvent)
	server.PATCH("/events/:id", controllers.UpdateEvent)
	server.DELETE("/events/:id", controllers.DeleteEvent)
}

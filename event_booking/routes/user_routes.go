package routes

import (
	"event-booking/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("/register", controllers.RegisterUser)
	server.POST("/login", controllers.LoginUser)
}

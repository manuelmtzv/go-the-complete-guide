package main

import (
	"event-booking/database"
	"event-booking/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":3000")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		fmt.Println(err)
		return
	}

	event.Id = 1
	event.UserId = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}

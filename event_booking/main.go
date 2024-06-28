package main

import (
	"event-booking/database"
	"event-booking/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/events", createEvent)

	server.Run(":3000")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events.",
		})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The provided Id must be a number",
		})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to retrieve the request entity.",
		})
		return
	}

	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Event with provided id (%v) was not found.", id),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"event": event,
	})
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

	err = event.Save()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create the requested event.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}

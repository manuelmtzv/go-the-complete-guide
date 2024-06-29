package controllers

import (
	"event-booking/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events.",
		})
		return
	}

	context.JSON(http.StatusOK, events)
}

func GetRawEventById(context *gin.Context, id int64) *models.Event {
	event, err := models.GetEventById(id)

	if err != nil {
		errMessage := "Unable to retrieve the event."
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errMessage,
		})
		return nil
	}

	if event == nil {
		errMessage := fmt.Sprintf("Event with provided id (%v) was not found.", id)
		context.JSON(http.StatusNotFound, gin.H{
			"message": errMessage,
		})
		return nil
	}

	return event
}

func GetEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The provided Id must be a number.	",
		})
		return
	}

	if event := GetRawEventById(context, id); event != nil {
		context.JSON(http.StatusOK, gin.H{
			"event": event,
		})
	}
}

func CreateEvent(context *gin.Context) {
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
		"message": "Event created.",
		"event":   event,
	})
}

func UpdateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The provided Id must be a number.",
		})
		return
	}

	event := GetRawEventById(context, id)

	if event == nil {
		return
	}

	var updatedFields map[string]interface{}
	if err := context.ShouldBindJSON(&updatedFields); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}

	if name, ok := updatedFields["name"].(string); ok {
		event.Name = name
	}
	if description, ok := updatedFields["description"].(string); ok {
		event.Description = description
	}
	if location, ok := updatedFields["location"].(string); ok {
		event.Location = location
	}
	if dateTimeStr, ok := updatedFields["datetime"].(string); ok {
		dateTime, err := time.Parse(time.RFC3339, dateTimeStr)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid dateTime format. Use RFC3339 format.",
			})
			return
		}
		event.DateTime = dateTime
	}

	if err := event.Update(); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update the requested event.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated.",
		"event":   event,
	})
}

func DeleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The provided Id must be a number.",
		})
		return
	}

	event := GetRawEventById(context, id)

	if event == nil {
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while trying to delete event.",
		})
		return
	}

	context.Status(http.StatusNoContent)
}

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

func FetchEvent(context *gin.Context, id int64) *models.Event {
	event, err := models.GetEventById(id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to retrieve the requested event.",
		})
		return nil
	}

	if event == nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Event with provided id (%v) was not found.", id),
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

	if event := FetchEvent(context, id); event != nil {
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
		return
	}

	event.UserId = context.GetInt64("userId")

	err = event.Save()

	if err != nil {
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
	userId := context.GetInt64("userId")
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The provided Id must be a number.",
		})
		return
	}

	event := FetchEvent(context, id)

	if event == nil {
		return
	}

	if validOwnership := event.ValidateOwnership(userId); !validOwnership {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "You are not authorized to update this event.",
		})
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
	userId := context.GetInt64("userId")
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The provided Id must be a number.",
		})
		return
	}

	event := FetchEvent(context, id)

	if event == nil {
		return
	}

	if validOwnership := event.ValidateOwnership(userId); !validOwnership {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "You are not authorized to delete this event.",
		})
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

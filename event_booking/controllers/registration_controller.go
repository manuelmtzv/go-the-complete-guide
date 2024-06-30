package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The provided Id must be a number.",
		})
		return
	}

	event := FetchEvent(context, eventId)

	if event == nil {
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create the registration to the event.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Your registration was set successfully.",
	})
}

func CancelEventRegistration() {

}

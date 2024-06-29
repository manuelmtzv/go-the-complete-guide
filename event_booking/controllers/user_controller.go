package controllers

import (
	"event-booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	user := models.User{}

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}

	foundUser, err := models.GetUser(user.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not validate the user's uniqueness.",
		})
		return
	}

	if foundUser != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The email is already been taken.",
		})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create the requested user.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully.",
	})
}

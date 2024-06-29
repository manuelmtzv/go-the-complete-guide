package controllers

import (
	"event-booking/models"
	"event-booking/utility"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchUser[T int64 | string](context *gin.Context, identifier T) *models.User {
	user, err := models.GetUser(identifier)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to retrieve the requested user.",
		})
	}

	if user == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("User with provided identifier (%v) was not found", identifier),
		})
	}

	return user
}

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

func LoginUser(context *gin.Context) {
	user := models.User{}
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Payload should contain 'email' and 'password'",
		})
		return
	}

	validCredentials, err := user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("User with provided email (%v) was not found.", user.Email),
		})
		return
	}

	if !validCredentials {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Email or password are incorrect.",
		})
		return
	}

	token, err := utility.GenerateJWT(user.Id, user.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate access token.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully authenticated!",
		"token":   token,
	})
}

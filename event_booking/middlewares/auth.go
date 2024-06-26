package middlewares

import (
	"event-booking/utility"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	var token string
	if bearerToken := strings.Split(context.Request.Header.Get("Authorization"), " ")[1:]; len(bearerToken) > 0 {
		token = bearerToken[0]
	}

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized.",
		})
		return
	}

	userId, err := utility.VerifyJwt(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or expired token.",
		})
		return
	}

	context.Set("userId", userId)

	context.Next()
}

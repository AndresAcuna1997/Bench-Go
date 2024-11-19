package middlewares

import (
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	authToken := c.Request.Header.Get("Authorization")

	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "A Token is required"})
		return
	}

	userId, err := utils.VerifyToken(authToken)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Could not verify token"})
		return
	}

	c.Set("userId", userId)

	c.Next()
}

package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse string to Int"})
		return
	}

	event, err := models.GetEventFromId(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user to the event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})

}

func CancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse string to Int"})
		return
	}

	var event models.Event

	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel the registratioon"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration canceled"})
}

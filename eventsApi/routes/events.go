package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetEventsHandler(c *gin.Context) {
	events, err := models.GetAllEvent()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not GET all events"})
		return
	}

	c.JSON(http.StatusOK, events)
}

func GetSingleEventHandler(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse string to Int"})
		return
	}

	event, err := models.GetEventFromId(eventId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"event": event})
}

func CreateEventHandler(c *gin.Context) {

	var newEvent models.Event

	err := c.ShouldBindJSON(&newEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data, there is a missing required field"})
		return
	}

	userId := c.GetInt64("userId")
	newEvent.UserID = userId

	err = newEvent.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": newEvent})
}

func UpdateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse string to Int"})
		return
	}

	event, err := models.GetEventFromId(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the spcified event"})
		return
	}

	userId := c.GetInt64("userId")

	if event.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"message": "Only the owner can modify this events"})
		return
	}

	var updatedEvent models.Event

	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data, there is a missing required field"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Updated the event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event updated", "event": updatedEvent})
}

func DeleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse string to Int"})
		return
	}

	event, err := models.GetEventFromId(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the spcified event"})
		return
	}

	userId := c.GetInt64("userId")

	if event.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"message": "Only the owner can delete this events"})
		return
	}

	err = event.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event deleted"})
}

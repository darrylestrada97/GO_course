package routes

import (
	"example/rest_api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEvent(c *gin.Context) {
	id := c.Param("id")
	event, err := models.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not fetch event",
		})
		return
	}
	c.JSON(http.StatusOK, event)
}
func CreateEvents(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad_request"})

		fmt.Println(err)
		return
	}
	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could_not_save_event",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "created",
		"message": event,
	})
}
func GetEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal_server_error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": events,
	})
}

func UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	event, err := models.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch event",
		})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad_request",
		})
		return
	}

	updatedEvent.ID = event.ID
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch event",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "updated",
	})
}

func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	event, err := models.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch event",
		})
		return
	}
	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not delete event",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})

}

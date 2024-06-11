package routes

import (
	"example/rest_api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterForEvent(c *gin.Context) {
	userId := c.GetInt("userId")
	id := c.Param("id")

	event, err := models.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch event",
		})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not register for event",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "registered",
	})
}

func cancelRegistration(c *gin.Context) {

	userId := c.GetInt("userId")
	id := c.Param("id")

	event, err := models.GetEvent(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch event",
		})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not cancel registration",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "cancelled",
	})
}

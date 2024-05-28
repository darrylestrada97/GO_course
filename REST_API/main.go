package main

import (
	"example/rest_api/db"
	"example/rest_api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := server.Run(":8080")
	if err != nil {
		println("Error: ", err)
		return
	}
}

func createEvents(c *gin.Context) {
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

	event.Save()
	c.JSON(http.StatusCreated, gin.H{
		"status":  "created",
		"message": event,
	})
}
func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, gin.H{
		"result": events,
	})
}
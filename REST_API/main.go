package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "events",
		})

	})
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

package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)
	server.POST("/events", CreateEvents)
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.PUT("/events/:id", UpdateEvent)
}

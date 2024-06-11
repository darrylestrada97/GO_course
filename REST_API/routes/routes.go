package routes

import (
	"example/rest_api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", CreateEvents)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)

}

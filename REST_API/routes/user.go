package routes

import (
	"example/rest_api/models"
	"example/rest_api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(c *gin.Context) {
	// signup a new user
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad_request"})
		fmt.Println(err)
		return
	}

	err = user.Save()
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "could not save user",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "created",
		"message": user,
	})
}

func login(c *gin.Context) {
	// login a user
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad_request"})
		fmt.Println(err)
		return
	}
	err = user.VerifyUser()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": err,
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate token",
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "logged in",
		"token":   token,
	})
}

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yamess/auth/models"
	"github.com/yamess/auth/utils"
	"log"
	"net/http"
)

func LoginUser(c *gin.Context) {
	var userCredentials models.UserCredentials
	var user models.User

	userCredentials = c.Keys["userCredentials"].(models.UserCredentials)

	user = models.User{Email: userCredentials.Email}
	result := user.GetRecord("email", user.Email)
	if result.Error != nil {
		logString := fmt.Sprintf("Enable to get data from database.\n%s", result.Error)
		log.Println(logString)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	if result.RowsAffected == 0 {
		log.Println("No User found. Please register first")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "No User found. Please register first"})
		return
	}

	isCorrect := utils.VerifyPassword(user.Password, userCredentials.Password)
	if !isCorrect {
		log.Println("User unauthorized")
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	tokenString, err := utils.GenerateToken(user.Id, user.Email)
	if err != nil {
		log.Println("Enable to generate token")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": models.Token{Token: tokenString}})
}

func LogoutUser(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {
		log.Println("Enable to parse token")
	}
}

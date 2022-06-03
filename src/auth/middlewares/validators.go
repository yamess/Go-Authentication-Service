package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yamess/auth/models"
	"log"
	"net/http"
)

func CredentialsValidator(c *gin.Context) {
	var userCredentials models.UserCredentials

	err := c.BindJSON(&userCredentials)
	if err != nil {
		logString := fmt.Sprintf("Error while parsing the request json data.\n%s", err.Error())
		log.Println(logString)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": logString})
		return
	}

	err = userCredentials.Validate()
	if err != nil {
		logString := fmt.Sprintf("Invalid data schema. %s", err.Error())
		log.Println(logString)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": logString})
		return
	}
	c.Set("userCredentials", userCredentials)
	c.Next()
}

func UserValidator(c *gin.Context) {
	var us models.User

	err := c.BindJSON(&us)
	if err != nil {
		logString := fmt.Sprintf("Error while parsing the request json data. %s", err.Error())
		log.Println(logString)
		c.JSON(http.StatusBadRequest, "Error while parsing the request data")
		return
	}

	err = us.Validate()
	if err != nil {
		logString := fmt.Sprintf("Invalid data schema. %s", err.Error())
		log.Println(logString)
		c.JSON(http.StatusBadRequest, gin.H{"error": logString})
		return
	}

	c.Set("authUser", us)
	c.Next()
}

func PasswordResetValidator(c *gin.Context) {
	var pwd models.UpdatePassword

	err := c.BindJSON(&pwd)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Enable to process request"})
		return
	}

	// Validation
	if err = pwd.Validate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad data schema"})
		return
	}

	c.Set("pwdData", pwd)
	c.Next()
}

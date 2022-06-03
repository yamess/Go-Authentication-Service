package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yamess/auth/models"
	"github.com/yamess/auth/utils"
	"log"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user models.User
	user = c.Keys["authUser"].(models.User)

	result := user.CreateRecord()
	if result.Error != nil {
		log.Printf("Unable to create record.\n%s", result.Error)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": "user created"})
}

func UpdatePassword(c *gin.Context) {
	var user models.User
	var pwd models.UpdatePassword

	id := uuid.Must(uuid.Parse(c.Param("id")))

	// Get password data and convert to JSON data type
	pwd = c.Keys["pwdData"].(models.UpdatePassword)
	data, er := json.Marshal(&pwd)
	if er != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Enable to process request"})
		return
	}

	// Check if old password is correct
	claims := c.Keys["claims"].(jwt.MapClaims)
	err := json.Unmarshal([]byte(string(data)), &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Email = claims["email"].(string)
	user.Id = id

	result := user.GetRecord("email", user.Email)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error})
		return
	}

	ok := utils.VerifyPassword(user.Password, pwd.OldPassword)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong Password"})
		return
	}

	// Update the password
	userId := uuid.Must(uuid.Parse(claims["user_id"].(string)))
	newPwd, e := utils.HashPassword(pwd.Password)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unable to process password"})
		return
	}

	user.Password = newPwd
	result = user.UpdatePassword(id, userId)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated with success"})
}

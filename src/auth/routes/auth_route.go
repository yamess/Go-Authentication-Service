package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yamess/auth/controllers"
	"github.com/yamess/auth/middlewares"
)

func AuthRoutes(a *gin.RouterGroup) {
	auth := a.Group("/auth")
	{
		auth.POST("/login", middlewares.CredentialsValidator, controllers.LoginUser)
		auth.POST("/logout", middlewares.Authenticate)
	}
}

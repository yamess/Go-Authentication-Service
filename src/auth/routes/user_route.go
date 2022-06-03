package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yamess/auth/controllers"
	"github.com/yamess/auth/middlewares"
)

func UserRoute(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/register", middlewares.UserValidator, controllers.CreateUser)
		user.GET("/me", middlewares.Authenticate)
		user.PATCH(
			"/reset-password/:id",
			middlewares.PasswordResetValidator, middlewares.Authenticate, controllers.UpdatePassword,
		)
	}
}

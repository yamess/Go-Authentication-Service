package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/yamess/auth/configs"
	"github.com/yamess/auth/routes"
)

func Run() {
	baseRoute := gin.Default()
	basePath := fmt.Sprintf("/api/v%s", configs.Version)
	v1 := baseRoute.Group(basePath)

	//Apply auth routes
	routes.AuthRoutes(v1)
	// Apply user route
	routes.UserRoute(v1)

	// Setting swagger doc endpoint
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run server
	err := baseRoute.Run(configs.Host)
	if err != nil {
		panic(err)
	}
}

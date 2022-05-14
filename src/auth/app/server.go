package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/yamess/auth/configs"
)

func Run() {
	baseRoute := gin.Default()
	basePath := fmt.Sprintf("/api/v%s", configs.Version)
	v1 := baseRoute.Group(basePath)

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := baseRoute.Run(configs.Host)
	if err != nil {
		panic(err)
	}
}

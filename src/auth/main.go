package main

import (
	"github.com/yamess/auth/app"
	"github.com/yamess/auth/configs"
	"github.com/yamess/auth/database"
	"github.com/yamess/auth/models"
)

func main() {
	// @title User API documentation
	// @version 1.0.0
	// @description     This is a simple rest api for category
	// @contact.name   API Support
	// @contact.url    http://www.swagger.io/support
	// @contact.email  support@swagger.io
	// @host localhost:8082
	// @BasePath /api/v1

	// Load environment variables
	configs.InitEnv()

	// Apply auto migration of the models
	database.Automigrate(models.User{})

	app.Run()
}

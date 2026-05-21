package main

import (
	"secret/config"
	"secret/middleware"
	"secret/routes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	gin.SetMode(config.GetConfig().ServerMode)

	router := gin.New()

	router.Use(middleware.Logrus(log))
	router.Use(gin.Recovery())
	routes.RegisterRoutes(router)

	port := config.GetConfig().ServerPort
	log.Info("Server starting on port ", port)
	router.Run(":" + port)
}

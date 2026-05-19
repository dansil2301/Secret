package main

import (
	"log"
	"secret/config"
	"secret/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.GetConfig().ServerMode)

	router := gin.Default()

	routes.RegisterRoutes(router)

	port := config.GetConfig().ServerPort
	log.Printf("Server starting on port %s", port)
	router.Run(":" + port)
}

package main

import (
	"log"
	"secret/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterRoutes(router)

	log.Printf("Server starting")
	router.Run(":8080")
}

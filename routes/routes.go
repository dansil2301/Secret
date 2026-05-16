package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	// Public
	public := router.Group("/api")
	{
		// Authentication routes
		public.POST("/login", func(c *gin.Context) {})
		public.POST("/register", func(c *gin.Context) {})

		public.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}

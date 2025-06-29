package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Define your routes here
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Add more routes as needed
	// r.POST("/example", exampleHandler)
	// r.GET("/users/:id", getUserHandler)
	// etc.
}

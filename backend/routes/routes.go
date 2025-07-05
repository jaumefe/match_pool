package routes

import (
	"match_pool_back/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	// Define your routes here
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/teams", controllers.GetTeams)
	r.GET("/scorers", controllers.GetScorers)
}

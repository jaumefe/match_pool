package routes

import (
	"match_pool_back/controllers"
	"match_pool_back/middleware"

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
	r.GET("/teams/:groupName", controllers.GetTeamsByGroup)
	r.GET("/scorers", controllers.GetScorers)
	r.GET("/scorers/:groupName/*position", controllers.GetScorersByGroupAndPosition)
	r.GET("/pool/teams", middleware.JWTMiddleware(), controllers.GetTeamsByUser)

	r.POST("/users", middleware.JWTMiddleware(), controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.POST("/pool/teams", middleware.JWTMiddleware(), controllers.SubmitTeamsUser)
}

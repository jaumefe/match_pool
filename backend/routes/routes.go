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
	r.GET("/scorers/teams/:teamId", controllers.GetScorerByTeamId)
	r.GET("/pool/teams", middleware.JWTMiddleware(), controllers.GetTeamsByUser)
	r.GET("/pool/scorers", middleware.JWTMiddleware(), controllers.GetScorerByUser)
	r.GET("/pool/points", middleware.JWTMiddleware(), controllers.GetPointsAllUsers)
	r.GET("/stages", controllers.GetStages)

	r.POST("/users", middleware.JWTMiddleware(), controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.POST("/pool/teams", middleware.JWTMiddleware(), controllers.SubmitTeamsUser)
	r.POST("/pool/scorers", middleware.JWTMiddleware(), controllers.SubmitScorersUser)
	r.POST("/matches", middleware.JWTMiddleware(), controllers.RegisterMatch)
	r.POST("/matches/id", middleware.JWTMiddleware(), controllers.RegisterMatchById)
	r.POST("/goals", middleware.JWTMiddleware(), controllers.RegisterGoals)
	r.POST("/goals/id", middleware.JWTMiddleware(), controllers.RegisterGoalsByPlayerID)
	r.POST("/match/id", middleware.JWTMiddleware(), controllers.GetMatchId)
	r.POST("/teams/pool_position", middleware.JWTMiddleware(), controllers.SetPoolPosition)
}

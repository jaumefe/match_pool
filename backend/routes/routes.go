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
	r.GET("/stages", controllers.GetStagesName)
	r.GET("/stages/points", controllers.GetStagesNameAndPoints)
	r.GET("/users/name", middleware.JWTMiddleware(), controllers.GetUserName)
	r.GET("/goals/stage", middleware.JWTMiddleware(), middleware.IsAdminMiddleware(), controllers.GetPointsGoalPerStage)

	r.POST("/users", middleware.JWTMiddleware(), middleware.IsAdminMiddleware(), controllers.RegisterUser)
	r.POST("/users/name", middleware.JWTMiddleware(), controllers.UpdateName)
	r.POST("/users/pass", middleware.JWTMiddleware(), controllers.UpdateUserPass)
	r.POST("/login", controllers.LoginUser)
	r.POST("/pool/teams", middleware.JWTMiddleware(), controllers.SubmitTeamsUser)
	r.POST("/pool/scorers", middleware.JWTMiddleware(), controllers.SubmitScorersUser)
	r.POST("/matches", middleware.JWTMiddleware(), middleware.IsAdminMiddleware(), controllers.RegisterMatch)
	r.POST("/matches/id", middleware.JWTMiddleware(), middleware.IsAdminMiddleware(), controllers.RegisterMatchById)
	r.POST("/goals", middleware.JWTMiddleware(), middleware.IsAdminMiddleware(), controllers.RegisterGoals)
	r.POST("/goals/id", middleware.JWTMiddleware(), middleware.IsAdminMiddleware(), controllers.RegisterGoalsByPlayerID)
	r.POST("/match/id", middleware.JWTMiddleware(), middleware.IsAdminMiddleware(), controllers.GetMatchId)
	r.POST("/teams/pool_position", middleware.JWTMiddleware(), middleware.IsAdminMiddleware(), controllers.SetPoolPosition)
	r.POST("/register/teams", middleware.JWTMiddleware(), middleware.IsAdminMiddleware(), controllers.RegisterTeamsFromCsv)
	r.POST("/register/scorers", middleware.JWTMiddleware(), middleware.IsAdminMiddleware(), controllers.RegisterScorersFromCsv)
	r.POST("/configuration/stage", middleware.JWTMiddleware(), middleware.IsAdminMiddleware(), controllers.SetStagePointsWin)
	r.POST("/configuration/goals", middleware.JWTMiddleware(), middleware.IsAdminMiddleware(), controllers.SetPointsGoalPerStage)
}

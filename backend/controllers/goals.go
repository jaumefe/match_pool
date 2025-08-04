package controllers

import (
	db "match_pool_back/database"
	"match_pool_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterGoals(c *gin.Context) {
	role, ok := c.Get("role")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found in context"})
		return
	}

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to register matches"})
		return
	}

	var input models.Goal
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.DB.Exec(PRAGMA_FOREIGN_KEYS_ON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enable foreign keys"})
		return
	}

	_, err = db.DB.Exec(REGISTER_GOAL_BY_PLAYER_NAME_AND_MATCH_ID, input.Player, input.MatchID, input.Goals)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register goal(s)"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Goal(s) registered successfully"})
}

func RegisterGoalsByPlayerID(c *gin.Context) {
	role, ok := c.Get("role")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found in context"})
		return
	}

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to register matches"})
		return
	}

	var input models.GoalByPlayerID
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.DB.Exec(PRAGMA_FOREIGN_KEYS_ON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enable foreign keys"})
		return
	}

	_, err = db.DB.Exec(REGISTER_GOAL_BY_PLAYER_AND_MATCH_ID, input.PlayerID, input.MatchID, input.Goals)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register goal(s)"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Goal(s) registered successfully"})
}

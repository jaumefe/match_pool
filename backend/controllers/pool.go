package controllers

import (
	"fmt"
	db "match_pool_back/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SubmitTeamsUser(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	var input struct {
		userID  int
		TeamsID []int `json:"teams_id"`
	}
	input.userID = userID.(int)

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input. %v", err)})
		return
	}

	for _, teamID := range input.TeamsID {
		_, err := db.DB.Exec("INSERT INTO user_teams (user_id, team_id) VALUES (?, ?)", input.userID, teamID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign team"})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Teams assigned successfully"})
}

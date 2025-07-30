package controllers

import (
	db "match_pool_back/database"
	"match_pool_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterMatch(c *gin.Context) {
	role, ok := c.Get("role")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found in context"})
		return
	}

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to register matches"})
		return
	}

	var input models.Match
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.Leg != 2 {
		input.Leg = 1
	}

	_, err := db.DB.Exec(PRAGMA_FOREIGN_KEYS_ON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enable foreign keys"})
		return
	}

	_, err = db.DB.Exec(REGISTER_MATCH_BY_TEAM_NAMES, input.TeamHome, input.TeamAway, input.Stage, input.PenaltyWinner, input.Leg, input.TeamHomeScore, input.TeamAwayScore)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register match"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Match registered successfully"})
}

func RegisterMatchById(c *gin.Context) {
	role, ok := c.Get("role")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found in context"})
		return
	}

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to register matches"})
		return
	}

	var input models.MatchById
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.Leg != 2 {
		input.Leg = 1
	}

	_, err := db.DB.Exec(REGISTER_MATCH_BY_ID, input.Leg, input.StageID, input.TeamHomeID, input.TeamAwayID, input.TeamHomeScore, input.TeamAwayScore, input.PenaltyWinnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register match"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Match registered successfully"})
}

package controllers

import (
	db "match_pool_back/database"
	"match_pool_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTeams(c *gin.Context) {
	rows, err := db.DB.Query("SELECT * FROM teams")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query teams"})
		return
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		if err := rows.Scan(&team.ID, &team.Name, &team.GroupName, &team.Value, &team.EliminatedAt, &team.PoolGroup); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan team"})
			return
		}
		teams = append(teams, team)
	}

	c.JSON(http.StatusOK, teams)
}

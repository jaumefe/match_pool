package controllers

import (
	db "match_pool_back/database"
	"match_pool_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTeams(c *gin.Context) {
	rows, err := db.DB.Query(TEAMS_QUERY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query teams"})
		return
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		if err := rows.Scan(&team.ID, &team.Name, &team.GroupName, &team.Value, &team.PoolPostion, &team.PoolGroup); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan team"})
			return
		}
		teams = append(teams, team)
	}

	c.JSON(http.StatusOK, teams)
}

func GetTeamsByGroup(c *gin.Context) {
	groupName := c.Param("groupName")

	rows, err := db.DB.Query(TEAMS_BY_GROUP_QUERY, groupName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query teams"})
		return
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		if err := rows.Scan(&team.ID, &team.Name, &team.GroupName, &team.Value, &team.PoolPostion, &team.PoolGroup); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan team"})
			return
		}
		teams = append(teams, team)
	}

	c.JSON(http.StatusOK, teams)
}

func SetPoolPosition(c *gin.Context) {
	role, ok := c.Get("role")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found in context"})
		return
	}

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to register matches"})
		return
	}

	var input models.Team
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.DB.Exec(SET_POOL_POSITION_TEAM, input.PoolPostion, input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pool position"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pool position updated successfully"})
}

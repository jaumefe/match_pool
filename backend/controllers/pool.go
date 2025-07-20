package controllers

import (
	"fmt"
	db "match_pool_back/database"
	"match_pool_back/models"
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

	if err := checkUniqueTeamsID(input.TeamsID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	points, err := checkUserTeamsPoints(input.TeamsID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to calculate points: %v", err)})
		return
	}

	if points > MAX_POINTS {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("User points: %d exceed maximum: %d", points, MAX_POINTS)})
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

func checkUserTeamsPoints(ids []int) (int, error) {
	_, err := db.DB.Exec(PRAGMA_FOREIGN_KEYS_ON)
	if err != nil {
		return 0, err
	}

	var totalPoints int
	for _, id := range ids {
		var points int
		row := db.DB.QueryRow(GET_VALUES_BY_TEAM_ID_QUERY, id)
		if err := row.Scan(&points); err != nil {
			return 0, err
		}
		totalPoints += points
	}
	return totalPoints, nil
}

func checkUniqueTeamsID(teams []int) error {
	teamSet := make(map[int]struct{})
	for _, teamID := range teams {
		if _, exists := teamSet[teamID]; exists {
			return fmt.Errorf("duplicate team ID found: %d", teamID)
		}
		teamSet[teamID] = struct{}{}
	}
	return nil
}

func GetTeamsByUser(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	_, err := db.DB.Exec(PRAGMA_FOREIGN_KEYS_ON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, err := db.DB.Query(GET_USER_TEAMS_QUERY, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		if err := rows.Scan(&team.ID, &team.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		teams = append(teams, team)
	}

	c.JSON(http.StatusOK, teams)
}

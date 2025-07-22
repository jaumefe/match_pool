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

	if err := checkUniqueID(input.TeamsID); err != nil {
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

	wasUpdated, err := updateUserTeams(input.TeamsID, input.userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update user teams: %v", err)})
		return
	}

	if wasUpdated {
		c.JSON(http.StatusOK, gin.H{"message": "Teams updated successfully"})
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

func checkUniqueID(ids []int) error {
	idSet := make(map[int]struct{})
	for _, id := range ids {
		if _, exists := idSet[id]; exists {
			return fmt.Errorf("duplicate ID found: %d", id)
		}
		idSet[id] = struct{}{}
	}
	return nil
}

func updateUserTeams(ids []int, userID int) (bool, error) {
	_, err := db.DB.Exec(PRAGMA_FOREIGN_KEYS_ON)
	if err != nil {
		return false, err
	}

	rows, err := db.DB.Query(GET_USER_TEAM_ID_BY_USER_ID_QUERY, userID)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	var existingTeamIDs []int
	for rows.Next() {
		var teamID int
		if err := rows.Scan(&teamID); err != nil {
			return false, err
		}
		existingTeamIDs = append(existingTeamIDs, teamID)
	}

	if len(existingTeamIDs) == 0 {
		return false, nil
	}

	for i, teamID := range ids {
		_, err := db.DB.Exec(UPDATE_USER_TEAMS, teamID, userID, existingTeamIDs[i])
		if err != nil {
			return false, err
		}
	}

	return true, nil
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
		if err := rows.Scan(&team.ID, &team.Name, &team.Value); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		teams = append(teams, team)
	}

	c.JSON(http.StatusOK, teams)
}

func SubmitScorersUser(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	var input struct {
		userID    int
		ScorersID []int `json:"scorers_id"`
	}
	input.userID = userID.(int)

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input. %v", err)})
		return
	}

	if err := checkUniqueID(input.ScorersID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wasUpdated, err := updateUserScorers(input.ScorersID, input.userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update user scorers: %v", err)})
		return
	}

	if wasUpdated {
		c.JSON(http.StatusOK, gin.H{"message": "Scorers updated successfully"})
		return
	}

	for _, scorerID := range input.ScorersID {
		_, err := db.DB.Exec("INSERT INTO user_scorers (player_id, user_id) VALUES (?, ?)", scorerID, input.userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign scorer"})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Scorers assigned successfully"})
}

func updateUserScorers(ids []int, userID int) (bool, error) {
	_, err := db.DB.Exec(PRAGMA_FOREIGN_KEYS_ON)
	if err != nil {
		return false, err
	}

	rows, err := db.DB.Query(GET_USER_SCORERS_ID_BY_USER_ID_QUERY, userID)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	var existingScorerIDs []int
	for rows.Next() {
		var scorerID int
		if err := rows.Scan(&scorerID); err != nil {
			return false, err
		}
		existingScorerIDs = append(existingScorerIDs, scorerID)
	}

	if len(existingScorerIDs) == 0 {
		return false, nil
	}

	for i, scorerID := range ids {
		_, err := db.DB.Exec(UPDATE_USER_SCORERS, scorerID, userID, existingScorerIDs[i])
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func GetScorerByUser(c *gin.Context) {
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

	rows, err := db.DB.Query(GET_USER_SCORERS_QUERY, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var scorers []models.Scorer
	for rows.Next() {
		var scorer models.Scorer
		if err := rows.Scan(&scorer.ID, &scorer.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		scorers = append(scorers, scorer)
	}
	c.JSON(http.StatusOK, scorers)
}

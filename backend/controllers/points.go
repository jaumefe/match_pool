package controllers

import (
	"database/sql"
	"fmt"
	db "match_pool_back/database"
	"match_pool_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPointsAllUsers(c *gin.Context) {
	_, err := db.DB.Exec(PRAGMA_FOREIGN_KEYS_ON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enable foreign keys"})
		return
	}

	users := []models.User{}
	rows, err := db.DB.Query(USERS_QUERY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query users"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Score); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user"})
			return
		}
		users = append(users, user)
	}

	pointsUser := make(map[string]int, len(users))
	for _, user := range users {
		var teams []models.Team
		rows, err := db.DB.Query(GET_USER_TEAMS_QUERY_NO_VALUE, user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user teams"})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var team models.Team
			if err := rows.Scan(&team.ID, &team.Name); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan team"})
				return
			}
			teams = append(teams, team)
		}
		pointsUser[user.Name], err = computePointsTeams(teams)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to compute points"})
			return
		}
	}
	c.JSON(http.StatusOK, pointsUser)
}

func computePointsTeams(teams []models.Team) (int, error) {
	points := 0
	rows, err := db.DB.Query(POINT_PER_STAGE_QUERY)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	pointsPerStage := make(map[int]int)
	for rows.Next() {
		var stageID, pointsWin int
		if err := rows.Scan(&stageID, &pointsWin); err != nil {
			return 0, err
		}
		pointsPerStage[stageID] = pointsWin
	}

	for _, team := range teams {
		rows, err := db.DB.Query(GET_TEAMS_POINTS_QUERY, team.ID, team.ID)
		if err != nil {
			return 0, err
		}
		defer rows.Close()

		var input struct {
			stageID         int
			teamHomeScore   int
			teamAwayScore   int
			teamHomeID      int
			teamAwayID      int
			penaltyWinnerID sql.NullInt16
		}
		for rows.Next() {
			if err := rows.Scan(&input.stageID, &input.teamHomeScore, &input.teamAwayScore, &input.teamHomeID, &input.teamAwayID, &input.penaltyWinnerID); err != nil {
				return 0, err
			}

			if input.penaltyWinnerID.Valid && int(input.penaltyWinnerID.Int16) == team.ID {
				points += pointsPerStage[input.stageID]
			} else if input.teamHomeID == team.ID {
				if input.teamHomeScore > input.teamAwayScore {
					points += pointsPerStage[input.stageID]
				}
			} else if input.teamAwayID == team.ID {
				if input.teamHomeScore < input.teamAwayScore {
					points += pointsPerStage[input.stageID]
				}
			}
		}

	}
	return points, nil
}

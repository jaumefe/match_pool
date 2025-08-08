package controllers

import (
	"database/sql"
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

	pointsTeamsUser := make(map[string]int, len(users))
	pointsScorerUser := make(map[string]int, len(users))
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
		pointsTeamsUser[user.Name], err = computePointsTeams(teams)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to compute points"})
			return
		}
		rows.Close()

		var scorers []models.Scorer
		rows, err = db.DB.Query(GET_USER_SCORERS_QUERY, user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user teams"})
			return
		}

		for rows.Next() {
			var scorer models.Scorer
			if err := rows.Scan(&scorer.ID, &scorer.Name); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan team"})
				return
			}
			scorers = append(scorers, scorer)
		}
		pointsScorerUser[user.Name], err = computePointsGoal(scorers)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to compute points"})
			return
		}
		pointsUser[user.Name] = pointsTeamsUser[user.Name] + pointsScorerUser[user.Name]
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
		rows, err := db.DB.Query(GET_MATCH_RESULT_QUERY, team.ID, team.ID)
		if err != nil {
			return 0, err
		}

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
			} else if input.penaltyWinnerID.Valid && int(input.penaltyWinnerID.Int16) != team.ID {
				continue
			} else if input.teamHomeID == team.ID {
				if input.teamHomeScore > input.teamAwayScore {
					points += pointsPerStage[input.stageID]
				} else if input.teamHomeScore == input.teamAwayScore {
					points++
				}
			} else if input.teamAwayID == team.ID {
				if input.teamHomeScore < input.teamAwayScore {
					points += pointsPerStage[input.stageID]
				} else if input.teamHomeScore == input.teamAwayScore {
					points++
				}
			}
		}
		rows.Close()

		var pool_position sql.NullInt16
		row := db.DB.QueryRow(GET_POOL_POSITION_TEAM_ID, team.ID)
		if err := row.Scan(&pool_position); err != nil {
			return 0, err
		}

		if pool_position.Valid {
			switch pool_position.Int16 {
			case 1:
				points += 2
			case 2:
				points++
			}
		}
	}

	return points, nil
}

func computePointsGoal(scorers []models.Scorer) (int, error) {
	points := 0
	rows, err := db.DB.Query(POINTS_PER_GOAL_QUERY)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	pointsGoalPerStage := make(map[int]int)
	for rows.Next() {
		var pointsGoal, stageId int
		if err := rows.Scan(&pointsGoal, &stageId); err != nil {
			return 0, err
		}
		pointsGoalPerStage[stageId] = pointsGoal
	}

	for _, scorer := range scorers {
		rows, err := db.DB.Query(GET_POINTS_SCORER_QUERY, scorer.ID)
		if err != nil {
			return 0, err
		}
		defer rows.Close()

		var input struct {
			goals   int
			stageID int
		}
		for rows.Next() {
			if err := rows.Scan(&input.goals, &input.stageID); err != nil {
				return 0, err
			}
			points += input.goals * pointsGoalPerStage[input.stageID]
		}
	}

	return points, nil
}

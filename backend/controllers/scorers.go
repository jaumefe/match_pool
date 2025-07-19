package controllers

import (
	"database/sql"
	db "match_pool_back/database"
	"match_pool_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetScorers(c *gin.Context) {
	_, err := db.DB.Exec(PRAGMA_FOREIGN_KEYS_ON) // Ensure foreign keys are enabled
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enable foreign keys"})
		return
	}

	rows, err := db.DB.Query(SCORERS_QUERY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query scorers"})
		return
	}
	defer rows.Close()

	var scorers []models.Scorer
	for rows.Next() {
		var scorer models.Scorer
		if err := rows.Scan(&scorer.ID, &scorer.Name, &scorer.Position, &scorer.Team); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan scorer"})
			return
		}
		scorers = append(scorers, scorer)
	}

	c.JSON(http.StatusOK, scorers)
}

func GetScorersByGroupAndPosition(c *gin.Context) {
	groupName := c.Param("groupName")
	position := c.Param("position")

	_, err := db.DB.Exec(PRAGMA_FOREIGN_KEYS_ON) // Ensure foreign keys are enabled
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enable foreign keys"})
		return
	}

	var rows *sql.Rows
	if position == "any" {
		rows, err = db.DB.Query(SCORERS_BY_GROUP_QUERY, groupName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query scorers"})
			return
		}
	} else {
		rows, err = db.DB.Query(SCORERS_BY_GROUP_AND_POSITION_QUERY, groupName, position)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query scorers"})
			return
		}
	}
	defer rows.Close()

	var scorers []models.Scorer
	for rows.Next() {
		var scorer models.Scorer
		if err := rows.Scan(&scorer.ID, &scorer.Name, &scorer.Position, &scorer.Team); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan scorer"})
			return
		}
		scorers = append(scorers, scorer)
	}
	c.JSON(http.StatusOK, scorers)
}

package controllers

import (
	db "match_pool_back/database"
	"match_pool_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetScorers(c *gin.Context) {
	rows, err := db.DB.Query("SELECT * FROM scorers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query scorers"})
		return
	}
	defer rows.Close()

	var scorers []models.Scorer
	for rows.Next() {
		var scorer models.Scorer
		if err := rows.Scan(&scorer.ID, &scorer.Name, &scorer.TeamID, &scorer.Position); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan scorer"})
			return
		}
		scorers = append(scorers, scorer)
	}

	c.JSON(http.StatusOK, scorers)
}

package controllers

import (
	db "match_pool_back/database"
	"match_pool_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStagesName(c *gin.Context) {
	rows, err := db.DB.Query(GET_ID_NAME_STAGES_QUERY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query stages"})
		return
	}
	defer rows.Close()

	var stages []models.Stage
	for rows.Next() {
		var stage models.Stage
		if err := rows.Scan(&stage.ID, &stage.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan stage"})
			return
		}
		stages = append(stages, stage)
	}
	c.JSON(http.StatusOK, stages)
}

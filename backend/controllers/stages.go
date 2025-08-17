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

func GetStagesNameAndPoints(c *gin.Context) {
	rows, err := db.DB.Query(GET_STAGES_NAME_POINTS)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query stages"})
		return
	}
	defer rows.Close()

	var stages []models.Stage
	for rows.Next() {
		var stage models.Stage
		if err := rows.Scan(&stage.Name, &stage.PointsWin); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan stage"})
			return
		}
		stages = append(stages, stage)
	}
	c.JSON(http.StatusOK, stages)
}

func SetStagePointsWin(c *gin.Context) {
	role, ok := c.Get("role")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found in context"})
		return
	}

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to register matches"})
		return
	}

	var input models.Stage
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.DB.Exec(SET_STAGES_WIN_POINTS_STAGE, input.Name, input.PointsWin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add stage name and points"})
		return
	}
}

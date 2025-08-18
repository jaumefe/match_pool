package controllers

import (
	"fmt"
	db "match_pool_back/database"
	"match_pool_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		if err := rows.Scan(&stage.ID, &stage.Name, &stage.PointsWin); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan stage"})
			return
		}
		stages = append(stages, stage)
	}
	c.JSON(http.StatusOK, stages)
}

func SetStagePointsWin(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "stage win points updated successfully"})
}

func GetPointsGoalPerStage(c *gin.Context) {
	_, err := db.DB.Exec(PRAGMA_FOREIGN_KEYS_ON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enable foreign keys"})
		return
	}

	rows, err := db.DB.Query(GET_POINTS_GOAL_PER_STAGE)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query points goal per stage"})
		return
	}
	defer rows.Close()

	var points []models.GoalsPointsStageName
	for rows.Next() {
		var point models.GoalsPointsStageName
		if err := rows.Scan(&point.PointsPerGoal, &point.Stage); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan points"})
			return
		}
		points = append(points, point)
	}

	c.JSON(http.StatusOK, points)
}

func SetPointsGoalPerStage(c *gin.Context) {
	var input models.GoalsPoints

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.DB.Exec(SET_POINTS_GOAL_PER_STAGE, input.PointsPerGoal, input.StageID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set points per goal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "goal per points per stage updated successfully"})
}

package controllers

import (
	"encoding/csv"
	db "match_pool_back/database"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterTeamsFromCsv(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to open file"})
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to read file"})
		return
	}

	// Check that the data is properly formatted
	headers := records[0]
	titles := []string{"teams", "value", "group_name", "pool_group"}
	if len(headers) != len(titles) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "incorrect csv format"})
		return
	}

	for i, h := range headers {
		if titles[i] != strings.ToLower(h) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected header option"})
			return
		}
	}

	_, err = db.DB.Exec(PRAGMA_FOREIGN_KEYS_ON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enable foreign keys"})
		return
	}

	for _, record := range records[1:] {
		team := record[0]
		value, err := strconv.Atoi(record[1])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to parse value field"})
			return
		}
		group_name := record[2]
		pool_group := record[3]

		_, err = db.DB.Exec(INSERT_TEAMS_FROM_CSV, team, group_name, value, pool_group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to register the team"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teams registered"})
}

func RegisterScorersFromCsv(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to open file"})
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to read file"})
		return
	}

	// Check that the data is properly formatted
	headers := records[0]
	titles := []string{"name", "team", "position"}
	if len(headers) != len(titles) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "incorrect csv format"})
		return
	}

	for i, h := range headers {
		if titles[i] != strings.ToLower(h) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected header option"})
			return
		}
	}

	_, err = db.DB.Exec(PRAGMA_FOREIGN_KEYS_ON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enable foreign keys"})
		return
	}

	for _, record := range records[1:] {
		var teamId int
		name := record[0]
		team := record[1]
		position := strings.ToLower(record[2])

		row := db.DB.QueryRow(GET_TEAM_ID_BY_NAME, team)
		if err := row.Scan(&teamId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to get team ID"})
			return
		}

		_, err = db.DB.Exec(INSERT_SCORERS_FROM_CSV, name, teamId, position)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to register the team"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teams registered"})
}

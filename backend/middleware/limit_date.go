package middleware

import (
	"match_pool_back/controllers"
	db "match_pool_back/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LimitDateProvider interface {
	GetLimitDate() (string, error)
}

type LimitDate struct{}

func (LimitDate) GetLimitDate() (string, error) {
	row := db.DB.QueryRow(controllers.GET_LIMIT_DATE)
	var limitDate string
	if err := row.Scan(&limitDate); err != nil {
		return "", err
	}
	return limitDate, nil
}

func IsLimitDateOver(provider LimitDateProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		limitDate, err := provider.GetLimitDate()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get limit date"})
			return
		}

		if limitDate <= time.Now().Format(time.DateOnly) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Limit date is over"})
			return
		}

		c.Next()
	}
}

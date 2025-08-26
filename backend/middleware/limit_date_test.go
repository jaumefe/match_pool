package middleware_test

import (
	"errors"
	"match_pool_back/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type LimitDateMock struct {
	limitDate string
	err       error
}

func (m LimitDateMock) GetLimitDate() (string, error) {
	return m.limitDate, m.err
}

func TestIsLimitDateOver(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name          string
		limitDateMock LimitDateMock
		wantCode      int
		wantString    string
	}{
		{
			name:          "db error",
			limitDateMock: LimitDateMock{"", errors.New("db error")},
			wantCode:      http.StatusInternalServerError,
			wantString:    "Failed",
		},
		{
			name:          "limit date over",
			limitDateMock: LimitDateMock{time.Now().Add(-24 * time.Hour).Format(time.DateOnly), nil},
			wantCode:      http.StatusForbidden,
			wantString:    "Limit date is over",
		},
		{
			name:          "limit date valid",
			limitDateMock: LimitDateMock{time.Now().Add(24 * time.Hour).Format(time.DateOnly), nil},
			wantCode:      http.StatusOK,
			wantString:    "ok",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.Use(middleware.IsLimitDateOver(tt.limitDateMock))
			r.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "ok"})
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantString)
		})
	}
}

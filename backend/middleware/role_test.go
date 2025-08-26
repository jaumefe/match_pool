package middleware_test

import (
	"match_pool_back/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter(role interface{}, withRole bool) *gin.Engine {
	r := gin.New()
	if withRole {
		r.Use(func(c *gin.Context) {
			c.Set("role", role)
		})
	}
	r.Use(middleware.IsAdminMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
	return r
}

func TestIsAdminMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		role       interface{}
		withRole   bool
		wantCode   int
		wantString string
	}{
		{
			name:       "admin autorizado",
			role:       "admin",
			withRole:   true,
			wantCode:   http.StatusOK,
			wantString: "ok",
		},
		{
			name:       "user no autorizado",
			role:       "user",
			withRole:   true,
			wantCode:   http.StatusForbidden,
			wantString: "Not authorized",
		},
		{
			name:       "sin role en contexto",
			withRole:   false,
			wantCode:   http.StatusBadRequest,
			wantString: "role not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupRouter(tt.role, tt.withRole)

			req, _ := http.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantString)
		})
	}
}

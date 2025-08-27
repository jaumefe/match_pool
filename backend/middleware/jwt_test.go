package middleware_test

import (
	"match_pool_back/auth"
	"match_pool_back/middleware"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func generateToken(claims auth.Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString(auth.JwtKey)
	return tokenStr
}

func TestJWTMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	auth.JwtKey = []byte("test-secret-key")

	tests := []struct {
		name         string
		authHeader   string
		claims       *auth.Claims
		wantCode     int
		wantContains string
	}{
		{
			name:         "missing header",
			authHeader:   "",
			wantCode:     http.StatusUnauthorized,
			wantContains: "Missing or invalid token",
		},
		{
			name:         "invalid token",
			authHeader:   "Bearer invalid.token.here",
			wantCode:     http.StatusUnauthorized,
			wantContains: "Unauthorized",
		},
		{
			name: "valid token",
			claims: &auth.Claims{
				UserID:   123,
				Username: "tester",
				Role:     "admin",
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
				},
			},
			wantCode:     http.StatusOK,
			wantContains: "ok",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.Use(middleware.JWTMiddleware())
			r.GET("/protected", func(c *gin.Context) {
				userID, _ := c.Get("user_id")
				c.JSON(http.StatusOK, gin.H{
					"message":  "ok",
					"user_id":  userID,
					"username": c.GetString("username"),
					"role":     c.GetString("role"),
				})
			})

			authHeader := tt.authHeader
			if tt.claims != nil {
				authHeader = "Bearer " + generateToken(*tt.claims)
			}

			req, _ := http.NewRequest("GET", "/protected", nil)
			if authHeader != "" {
				req.Header.Set("Authorization", authHeader)
			}

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantContains)

			if tt.wantCode == http.StatusOK {
				assert.Contains(t, w.Body.String(), strconv.Itoa(tt.claims.UserID))
				assert.Contains(t, w.Body.String(), tt.claims.Username)
				assert.Contains(t, w.Body.String(), tt.claims.Role)
			}
		})
	}
}

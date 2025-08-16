package controllers

import (
	"match_pool_back/auth"
	db "match_pool_back/database"
	"match_pool_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	role, ok := c.Get("role")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found in context"})
		return
	}

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to register matches"})
		return
	}

	var input struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.Role == "" {
		input.Role = "user"
	}

	// Hash password to store it
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	_, err = db.DB.Exec("INSERT INTO users(name, token, role) VALUES (?, ?, ?)", input.Name, string(hash), input.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

func LoginUser(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var storedHash string
	var userID int
	var role string
	rows, err := db.DB.Query(USERS_TOKEN_BY_NAME_QUERY, input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&userID, &storedHash, &role); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user"})
			return
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	jwt, err := auth.GenerateJWT(userID, input.Name, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "bearer": jwt, "role": role})
}

func UpdateName(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found in context"})
		return
	}

	var input models.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.DB.Exec(UPDATE_USER_NAME, input.Name, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user name"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user name updated successfully"})
}

func GetUserName(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found in context"})
		return
	}

	var name string
	row := db.DB.QueryRow(GET_USER_NAME, id)
	if err := row.Scan(&name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to get user name"})
		return
	}

	c.JSON(http.StatusOK, name)
}

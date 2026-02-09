package controllers

import (
	"database/sql"
	"net/http"
	"task-manager-api/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// --------------------
// Register User
// --------------------
func Register(c *gin.Context) {

	var data struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Read JSON Body
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	// Hash Password
	hash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	// Insert User
	query := "INSERT INTO users(name,email,password) VALUES(?,?,?)"

	_, err := config.DB.Exec(query, data.Name, data.Email, hash)

	if err != nil {

		// Duplicate Email Check
		if err == sql.ErrNoRows {
			c.JSON(400, gin.H{"error": "User Exists"})
			return
		}

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Registration Successful"})
}

// --------------------
// Login User
// --------------------
func Login(c *gin.Context) {

	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Input"})
		return
	}

	var id int
	var hash string

	// Get User
	query := "SELECT id,password FROM users WHERE email=?"

	err := config.DB.QueryRow(query, data.Email).Scan(&id, &hash)

	if err != nil {

		c.JSON(401, gin.H{"error": "Invalid Email or Password"})
		return
	}

	// Compare Password
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(data.Password))

	if err != nil {

		c.JSON(401, gin.H{"error": "Invalid Email or Password"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login Successful",
		"user_id": id,
	})
}

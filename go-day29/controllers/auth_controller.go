package controllers

import (
	"task-manager-api/config"

	"fmt"

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

	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Input"})
		return
	}

	// Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		c.JSON(500, gin.H{"error": "Password Hash Failed"})
		return
	}

	// Insert User
	query := "INSERT INTO users(name,email,password) VALUES(?,?,?)"

	result, err := config.DB.Exec(query, data.Name, data.Email, string(hash))

	if err != nil {

		fmt.Println("❌ DB ERROR:", err) // Terminal-এ দেখাবে

		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, _ := result.LastInsertId()

	fmt.Println("✅ User Inserted ID:", id)

	c.JSON(200, gin.H{
		"message": "Registration Successful",
		"id":      id,
	})
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

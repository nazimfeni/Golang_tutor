package controllers

import (
	"task-manager-api/config"

	"github.com/gin-gonic/gin"
)

// Create Task
func CreateTask(c *gin.Context) {

	var data struct {
		Title string `json:"title"`
	}

	c.BindJSON(&data)

	query := "INSERT INTO tasks(user_id,title,status) VALUES(?,?,?)"

	_, err := config.DB.Exec(query, 1, data.Title, "pending")

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task Created"})
}

// Get Tasks
func GetTasks(c *gin.Context) {

	rows, _ := config.DB.Query("SELECT id,title,status FROM tasks")

	var tasks []map[string]interface{}

	for rows.Next() {

		var id int
		var title, status string

		rows.Scan(&id, &title, &status)

		tasks = append(tasks, gin.H{
			"id":     id,
			"title":  title,
			"status": status,
		})
	}

	c.JSON(200, tasks)
}

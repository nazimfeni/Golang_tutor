package routes

import (
	"task-manager-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")

	{
		api.POST("/tasks", controllers.CreateTask)
		api.GET("/tasks", controllers.GetTasks)
	}
}


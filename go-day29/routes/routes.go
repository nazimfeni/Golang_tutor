// package routes

// import (
// 	"task-manager-api/controllers"

// 	"github.com/gin-gonic/gin"
// )

// func SetupRoutes(r *gin.Engine) {

// 	api := r.Group("/api")

// 	{
// 		api.POST("/tasks", controllers.CreateTask)
// 		api.GET("/tasks", controllers.GetTasks)
// 	}
// }

package routes

import (
	"task-manager-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")

	{
		// Auth
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// Tasks
		api.POST("/tasks", controllers.CreateTask)
		api.GET("/tasks", controllers.GetTasks)
	}
}

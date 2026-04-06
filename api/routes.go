package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/RevDau-PallaviShinde/Go-Cli-App/docs" // swagger docs
)

// SetupRouter configures the Gin router with all API routes.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS — allows all origins (same as the manual version, but cleaner)
	r.Use(cors.Default())

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Task routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("/tasks", CreateTask)
		v1.GET("/tasks", GetTasks)
		v1.GET("/tasks/search", SearchTasks)
		v1.GET("/tasks/summary", GetSummary)
		v1.GET("/tasks/export", ExportTasks)
		v1.PUT("/tasks/:id", UpdateTask)
		v1.DELETE("/tasks/:id", DeleteTask)
	}

	return r
}

package api

import (
	"fmt"
	"log"
	"os"

	"github.com/RevDau-PallaviShinde/Go-Cli-App/docs"
	"github.com/joho/godotenv"
)

// StartServer launches the REST API using config from .env.
func StartServer() {
	// Load .env file (OK if it's missing)
	_ = godotenv.Load()

	ip := os.Getenv("SERVER_IP")
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// Always set Host to empty so it uses the current browser address (fixes CORS/Fetch errors)
	docs.SwaggerInfo.Host = ""

	router := SetupRouter()

	fmt.Println("========================================")
	fmt.Printf("  Task Manager API running at: %s:%s\n", ip, port)
	fmt.Printf("  Swagger UI: http://%s:%s/swagger/index.html\n", ip, port)
	fmt.Println("========================================")

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

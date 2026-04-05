package api

import (
	"fmt"
	"log"
)

// StartServer launches the REST API on port 8080.
func StartServer() {
	router := SetupRouter()

	fmt.Println("========================================")
	fmt.Println("  Task Manager API running on :8080")
	fmt.Println("  Swagger UI: http://13.233.237.253:8080/swagger/index.html")
	fmt.Println("========================================")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

package main

import (
	"log"
	"os"

	"github.com/SpaceTesla/orbis-be/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found. Using defaults.")
	}

	// Initialize Gin router
	r := gin.Default()

	// Enable CORS (allow frontend to connect)
	r.Use(cors.Default())

	// Define routes
	routes.SetupRoutes(r)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port
	}
	log.Println("🚀 Orbis API running on port", port)
	err = r.Run(":" + port)
	if err != nil {
		return
	}
}

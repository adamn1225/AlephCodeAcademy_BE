package main

import (
	"alephcode-backend/config"
	"alephcode-backend/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	if os.Getenv("GIN_MODE") != "release" {
		_ = godotenv.Load(".env.local")
	}

	// CORS setup
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	config.ConnectDB()
	routes.RegisterRoutes(r, config.DB)

	r.Run(":8080")
}

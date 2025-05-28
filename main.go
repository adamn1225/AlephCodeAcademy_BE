package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"alephcode-backend/config"
	"alephcode-backend/routes"
)

func main() {
	r := gin.Default()

	// CORS setup
	r.Use(cors.Default())

	config.ConnectDB()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}

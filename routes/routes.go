package routes

import (
	"alephcode-backend/handlers"
	"alephcode-backend/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// Missions
	r.GET("/api/missions", handlers.GetAllMissions)
	r.POST("/api/missions", handlers.SubmitMission)

	// Auth
	auth := r.Group("/api")
	{
		auth.POST("/login", handlers.LoginHandler(db))
		auth.POST("/parent/signup", handlers.ParentSignupHandler(db))

	}

	// Protected routes (e.g., RequireAuth middleware per role)
	parent := r.Group("/api/parent")
	parent.Use(middlewares.RequireAuth("parent"))
	{

	}

	student := r.Group("/api/student")
	student.Use(middlewares.RequireAuth("student"))
	{

	}

	teacher := r.Group("/api/teacher")
	teacher.Use(middlewares.RequireAuth("teacher"))
	{

	}

	admin := r.Group("/api/admin")
	admin.Use(middlewares.RequireAuth("admin"))
	{

	}
}

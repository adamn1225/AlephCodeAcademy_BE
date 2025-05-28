package routes

import (
    "github.com/gin-gonic/gin"
    "alephcode-backend/handlers"
)

func RegisterRoutes(r *gin.Engine) {
    r.GET("/api/missions", handlers.GetAllMissions)
    r.POST("/api/missions", handlers.SubmitMission)
}

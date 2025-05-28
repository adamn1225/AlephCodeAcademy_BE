package handlers

import (
    "net/http"
    "alephcode-backend/config"
    "alephcode-backend/models"
    "github.com/gin-gonic/gin"
)

func GetAllMissions(c *gin.Context) {
    var missions []models.Mission
    config.DB.Find(&missions)
    c.JSON(http.StatusOK, missions)
}

func SubmitMission(c *gin.Context) {
    var mission models.Mission
    if err := c.BindJSON(&mission); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&mission)
    c.JSON(http.StatusCreated, mission)
}

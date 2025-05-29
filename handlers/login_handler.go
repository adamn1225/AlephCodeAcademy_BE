package handlers

import (
	"alephcode-backend/middlewares"
	"alephcode-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		var teacher models.Teacher
		if err := db.Where("email = ?", body.Email).First(&teacher).Error; err == nil {
			if bcrypt.CompareHashAndPassword([]byte(teacher.PasswordHash), []byte(body.Password)) == nil {
				token, _ := middlewares.GenerateJWT("teacher", teacher.ID)
				c.JSON(http.StatusOK, gin.H{"token": token, "role": "teacher"})
				return
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

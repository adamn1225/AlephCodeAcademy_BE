// handlers/parents_auth.go

package handlers

import (
	"alephcode-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ParentSignupHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			FullName       string `json:"fullName" binding:"required"`
			Email          string `json:"email" binding:"required,email"`
			PhoneNumber    string `json:"phoneNumber" binding:"required"`
			Password       string `json:"password" binding:"required,min=6"`
			Description    string `json:"description"`
			ProfilePicture string `json:"profilePicture"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Check if user already exists
		var existing models.Parent
		if err := db.Where("email = ?", input.Email).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
			return
		}

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
			return
		}

		parent := models.Parent{
			FullName:       input.FullName,
			Email:          input.Email,
			PhoneNumber:    input.PhoneNumber,
			PasswordHash:   string(hashedPassword),
			Description:    input.Description,
			ProfilePicture: input.ProfilePicture,
		}

		if err := db.Create(&parent).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Parent created successfully"})
	}
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"ddx_hackathon_backend/helpers"
	"ddx_hackathon_backend/models"
)

func CreateUser(c *gin.Context, db *gorm.DB) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	role := "client"
	avatarUrl := helpers.GenerateAvatarURL(input.Email, role)
	user := models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  string(hashedPassword),
		Role:      role,
		AvatarUrl: &avatarUrl,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func LoginUser(c *gin.Context, db *gorm.DB) {
	var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
	}
	
	var user models.User
	if err := db.Preload("TrainerProfile.Specialties").Where("email = ?", input.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
	}
	
	// Проверка пароля
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
	}
	
	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context, db *gorm.DB) {
    var users []models.User
    if err := db.Find(&users).Error; err != nil {
        c.AbortWithStatus(http.StatusNotFound)
    } else {
        c.JSON(http.StatusOK, users)
    }
}

func DeleteUser(c *gin.Context, db *gorm.DB) {
    id := c.Params.ByName("id")
    var user models.User
    if err := db.Where("id = ?", id).First(&user).Error; err != nil {
        c.AbortWithStatus(http.StatusNotFound)
    } else {
        db.Delete(&user)
        c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
    }
}

func UpdateUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = string(hashedPassword)
	}
    user.AvatarUrl = input.AvatarUrl

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

package handlers

import (
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"haircare-backend/internal/database"
	"haircare-backend/internal/models"
)

var (
	hasUpperCase   = regexp.MustCompile(`[A-Z]`)
	hasNumber      = regexp.MustCompile(`[0-9]`)
	hasSpecialChar = regexp.MustCompile(`[!@#\$%\^&\*]`)
	minLength      = 8
)

func isValidPassword(password string) bool {
	return len(password) >= minLength &&
		hasUpperCase.MatchString(password) &&
		hasNumber.MatchString(password) &&
		hasSpecialChar.MatchString(password)
}

func Register(c *gin.Context) {
	var user models.Hairdresser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidPassword(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must have 8 chars, 1 uppercase, 1 number, 1 special char"})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Hairdresser registered successfully!"})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var user models.Hairdresser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

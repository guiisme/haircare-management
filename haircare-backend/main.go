package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"haircare-backend/internal/database"
	"haircare-backend/internal/handlers"
	"haircare-backend/internal/middleware"
	"haircare-backend/internal/models"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Hairdresser{}, &models.Client{})

	r := gin.Default()
	
	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	
	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/clients", handlers.CreateClient)
		protected.GET("/clients/:hairdresserID", handlers.ListClients)
		protected.PUT("/clients/:id", handlers.UpdateClient)
		protected.DELETE("/clients/:id", handlers.DeleteClient)
	}

	r.Run(":8080")
}

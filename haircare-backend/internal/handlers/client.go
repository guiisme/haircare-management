package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"haircare-backend/internal/database"
	"haircare-backend/internal/models"
)

func CreateClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Client added successfully"})
}

func ListClients(c *gin.Context) {
	hairdresserID := c.Param("hairdresserID")
	var clients []models.Client
	database.DB.Where("hairdresser_id = ?", hairdresserID).Find(&clients)
	fmt.Printf("Fetched clients: %+v\n", clients)
	c.JSON(http.StatusOK, clients)
	
}

func UpdateClient(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := database.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	var input models.Client
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&client).Updates(input)
	c.JSON(http.StatusOK, client)
}

func DeleteClient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Delete(&models.Client{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Client deleted"})
}

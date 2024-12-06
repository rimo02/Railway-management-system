package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"Railway-management-system/models"
)

func AddTrain(c *gin.Context, db *gorm.DB) {
	var train models.Train
	if err := c.ShouldBindJSON(&train); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := db.Create(&train).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add train"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Train added successfully"})
}

package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"Railway-management-system/models"
	"Railway-management-system/utils"
)

func BookSeat(c *gin.Context, db *gorm.DB) {
	userID, err := utils.GetUserIdFromJWTToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var req struct {
		TrainID uint `json:"train_id"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request"})
		return
	}

	err = db.Transaction(func(txn *gorm.DB) error {
		var train models.Train
		if err := db.Where("id = ?").First(&train).Error; err != nil {
			return err
		}

		if train.AvailableSeats < 1 {
			return errors.New("no Seats remaining")
		}
		train.AvailableSeats--
		if err := txn.Save(&train).Error; err != nil {
			return err
		}

		booking := models.Booking{
			UserID:     userID,
			TrainID:    train.ID,
			SeatNumber: train.TotalSeats - train.AvailableSeats,
		}

		if err := txn.Create(&booking).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seat booked successfully"})
}

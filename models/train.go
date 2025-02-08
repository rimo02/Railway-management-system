package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Train struct {
	ID             string `gorm:"type:char(36);primaryKey"`
	TrainID        int64  `gorm:"uniqueIndex; autoIncrement"`
	Source         string `json:"source"`
	Destination    string `json:"destination"`
	TotalSeats     int    `json:"total_seats"`
	AvailableSeats int    `json:"available_seats"`
}

// BeforeCreate hook to generate UUID and TrainID
func (train *Train) BeforeCreate(tx *gorm.DB) (err error) {
	train.ID = uuid.New().String()
	return
}

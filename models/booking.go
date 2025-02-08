package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	ID         string `gorm:"type:char(36);primaryKey"`
	UserID     int64  `json:"user_id"`
	TrainID    int64  `json:"train_id"`
	SeatNumber int    `json:"seat_no"`
}

func (booking *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	booking.ID = uuid.New().String()
	return
}

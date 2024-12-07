package models

type Booking struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint `json:"user_id"`
	TrainID    uint `json:"train_id"`
	SeatNumber int  `json:"seat_no"`
}

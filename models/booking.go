package models

type Booking struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	TrainID    uint
	SeatNumber int
}

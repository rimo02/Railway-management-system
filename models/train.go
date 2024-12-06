package models

type Train struct {
	ID             uint   `gorm:"primaryKey"`
	Source         string `json:"source"`
	Destination    string `json:"destination"`
	TotalSeats     int    `json:"total_seats"`
	AvailableSeats int    `json:"available_seats"`
}

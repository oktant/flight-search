package models

import "github.com/google/uuid"

type Flight struct {
	Id              uuid.UUID `gorm:"primary_key" json:"id"`
	FlightNo        string    `gorm:"flight_number" json:"flightNo"`
	Source          string    `gorm:"source" json:"source"`
	Destination     string    `gorm:"destination" json:"destination"`
	TravelStartDate string    `gorm:"travel_start_date" json:"travelStartDate"`
	TravelEndDate   string    `gorm:"travel_end_date" json:"travelEndDate"`
}

type Booking struct {
	Id       uuid.UUID `gorm:"primary_key" json:"id"`
	Username string    `json:"username"`
	FlightId string    `json:"flightId"`
	Flight   *Flight   `json:"flight,omitempty"`
}

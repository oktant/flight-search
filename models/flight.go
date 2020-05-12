package models

type Flight struct {
	Id              string `gorm:"primary_key" json:"id"`
	FlightNo        string `gorm:"flight_number" json:"flightNo"`
	Source          string `gorm:"source" json:"source"`
	Destination     string `gorm:"destination" json:"destination"`
	TravelStartDate string `gorm:"travel_start_date" json:"travelStartDate"`
	TravelEndDate   string `gorm:"travel_end_date" json:"travelEndDate"`
}

type Booking struct {
	Id string `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	FlightId string `json:"flightId"`
	Flight  *Flight `json:"flight,omitempty"`
}

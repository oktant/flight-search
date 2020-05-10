package models

type SearchResults struct {
	FlightNo        string `json: "flightNo"`
	Source          string `json: "source"`
	Destination     string `json: "destination"`
	TravelStartDate string `json: "travelStartDate"`
	TravelEndDate   string `json: "travelEndDate"`
}

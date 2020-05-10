package handlers

import (
	"net/http"
	"github.com/srinivasvinay/flight-search/models"
	"strconv"
	"encoding/json"
)

func GetFlights(w http.ResponseWriter, r *http.Request) {
	a := make([]models.SearchResults, 0)
	var i int64 = 0
	for ; i < 10; i++ {
		a = append(a, models.SearchResults{
			FlightNo:        "Flight-No-" + strconv.FormatInt(i, 10),
			Source:          r.URL.Query()["source"][0],
			Destination:     r.URL.Query()["dest"][0],
			TravelStartDate: r.URL.Query()["st-date"][0],
			TravelEndDate:   r.URL.Query()["ed-date"][0],
		})
	}
	output, _ := json.Marshal(&a)
	w.Write(output)
}

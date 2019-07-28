// taxi models for data returned from big query
package taxi

// average speed of all trips on a day
type averageSpeedDay struct {
	AverageSpeed float64 `json:"average_speed"`
}

// total number trips on a day
type totalTripsDay struct {
	Date       string
	TotalTrips string `json:"total_trips"`
}

// pick up location and fare price
type farePickupLocation struct {
	Lng  string `json:"pickup_longitude"`
	Lat  string `json:"pickup_latitude"`
	Fare string `json:"fare_amount"`
}

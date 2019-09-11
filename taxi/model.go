// Package taxi contains library for the taxi entity.
// Handler, model, query, repository, service, data.
package taxi

// Average speed of trips on a date.
type AverageSpeedByDay struct {
	AverageSpeed float64 `json:"average_speed"`
}

// Total number of trips on each date.
type TotalTripsByDay struct {
	Date       string `json:"date"`
	TotalTrips int    `json:"total_trips"`
}

// Pick up location by latitude and longitude and fare price.
type FarePickupByLocation struct {
	Lng  float64 `bigquery:"pickup_longitude"`
	Lat  float64 `bigquery:"pickup_latitude"`
	Fare float64 `bigquery:"fare_amount"`
}

// S2ID location and fare amount.
type S2idFare struct {
	S2id string  `json:"s2id"`
	Fare float64 `json:"fare"`
}
